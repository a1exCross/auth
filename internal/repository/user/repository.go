package userrepository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/a1exCross/auth/internal/model"
	"github.com/a1exCross/auth/internal/repository"

	"github.com/a1exCross/common/pkg/client/db"
	"github.com/a1exCross/common/pkg/filter"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

const (
	tableName = "users"

	idColumn        = "id"
	usernameColumn  = "username"
	nameColumn      = "name"
	emailColumn     = "email"
	roleColumn      = "role"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

// NewRepository - возвращает методы для работы с репозиторием пользователей
func NewRepository(db db.Client) repository.UserRepository {
	return repo{
		db: db,
	}
}

type repo struct {
	db db.Client
}

func (r repo) Create(ctx context.Context, params *model.UserCreate) (int64, error) {
	insertBuilder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(usernameColumn, nameColumn, emailColumn, roleColumn, passwordColumn).
		Values(params.Info.Username, params.Info.Name, params.Info.Email, params.Info.Role, params.Password).
		Suffix(fmt.Sprintf("RETURNING %s", idColumn))

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return 0, errors.Errorf("error at parse sql builder: %v", err)
	}

	var id int64

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, errors.Errorf("error at query to database: %v", err)
	}

	return id, nil
}

func (r repo) Get(ctx context.Context, filters filter.Filter) (*model.User, error) {
	selectBuilder := sq.Select(idColumn, usernameColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn, passwordColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName)

	for _, condition := range filters.Conditions {
		selectBuilder = selectBuilder.Where(sq.Eq{condition.Key: condition.Value})
	}

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, errors.Errorf("error at parse sql builder: %v", err)
	}

	q := db.Query{
		Name:     "user_repository.GetByID",
		QueryRaw: query,
	}

	var user model.User

	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Errorf("user not found")
		}
		return nil, errors.Errorf("error at query to database: %v", err)
	}

	return &user, nil
}

func (r repo) Delete(ctx context.Context, id int64) error {
	deleteBuilder := sq.Delete(tableName).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return errors.Errorf("error at parse sql builder: %v", err)
	}

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return errors.Errorf("error at query to database: %v", err)
	}

	return nil
}

func (r repo) Update(ctx context.Context, params *model.UserUpdate) error {
	updateBuilder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, params.Info.Name).
		Set(emailColumn, params.Info.Email).
		Set(roleColumn, params.Info.Role).
		Where(sq.Eq{idColumn: params.ID})

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return errors.Errorf("error at parse sql builder: %v", err)
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return errors.Errorf("error at query to database: %v", err)
	}

	return nil
}
