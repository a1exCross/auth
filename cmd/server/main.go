package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/a1exCross/auth/internal/config"
	pbUser "github.com/a1exCross/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type service struct {
	pbUser.UnimplementedUserV1Server
	pool *pgxpool.Pool
}

const table = "users"
const (
	id        = "id"
	name      = "name"
	email     = "email"
	role      = "role"
	password  = "password"
	createdAt = "created_at"
	updatedAt = "updated_at"
)

func (s service) Get(ctx context.Context, req *pbUser.GetRequest) (*pbUser.GetResponse, error) {
	selectBuilder := sq.Select(id, name, email, role, "created_at", "updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table).
		Where("id = ?", req.GetId())

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error at parse sql builder: %v", err)
	}

	row := s.pool.QueryRow(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error at query to database: %v", err)
	}

	var id int64
	var name, email string
	var role uint8
	var createdAt time.Time
	var updatedAt sql.NullTime

	err = row.Scan(&id, &name, &email, &role, &createdAt, &updatedAt)
	if err != nil {
		return nil, fmt.Errorf("error at scan row: %v", err)
	}

	var updatedAtTime *timestamppb.Timestamp
	if updatedAt.Valid {
		updatedAtTime = timestamppb.New(updatedAt.Time)
	}

	return &pbUser.GetResponse{
		User: &pbUser.User{
			Id: id,
			Info: &pbUser.UserInfo{
				Name:  name,
				Email: email,
				Role:  pbUser.UserRole(role),
			},
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: updatedAtTime,
		},
	}, nil
}

func (s service) Create(ctx context.Context, req *pbUser.CreateRequest) (*pbUser.CreateResponse, error) {
	if req.Pass.Password != req.Pass.PasswordConfirm {
		return nil, fmt.Errorf("passwords mismatch")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Pass.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed hash password: %v", err)
	}

	insertBuilder := sq.Insert(table).
		PlaceholderFormat(sq.Dollar).
		Columns(name, email, role, password).
		Values(req.Info.GetName(), req.Info.GetEmail(), req.Info.GetRole(), string(hashedPassword)).
		Suffix("RETURNING id")

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error at parse sql builder: %v", err)
	}

	var id int64

	err = s.pool.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("error at query to database: %v", err)
	}

	return &pbUser.CreateResponse{
		Id: id,
	}, nil
}

func (s service) Update(ctx context.Context, req *pbUser.UpdateRequest) (*empty.Empty, error) {
	updateBuilder := sq.Update(table).
		PlaceholderFormat(sq.Dollar).
		Set(name, req.GetName().Value).
		Set(email, req.GetEmail().Value).
		Set(role, req.GetRole()).
		Where("id = ?", req.GetId())

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error at parse sql builder: %v", err)
	}

	_, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error at query to database: %v", err)
	}

	return &empty.Empty{}, nil
}

func (s service) Delete(ctx context.Context, req *pbUser.DeleteRequest) (*empty.Empty, error) {
	deleteBuilder := sq.Delete(table).
		Where("id = ?", req.GetId()).
		PlaceholderFormat(sq.Dollar)

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error at parse sql builder: %v", err)
	}

	_, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error at query to database: %v", err)
	}

	return &empty.Empty{}, nil
}
func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load environments: %v", err)
	}

	grpcConf, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to create grpc config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConf.Address())
	if err != nil {
		log.Fatalf("failed to connect grpc server: %v", err)
	}

	log.Printf("Listen and serve at %s", grpcConf.Address())

	pgConf, err := config.NewPGConfig()

	log.Println(pgConf)

	if err != nil {
		log.Fatalf("failed to create pg config: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConf.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)

	pbUser.RegisterUserV1Server(
		s, service{
			pool: pool,
		})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc server: %v", err)
	}
}
