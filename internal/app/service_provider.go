package app

import (
	userAPI "github.com/a1exCross/auth/internal/api/user"
	"github.com/a1exCross/auth/internal/client/db"
	"github.com/a1exCross/auth/internal/client/db/pg"
	"github.com/a1exCross/auth/internal/client/db/transaction"
	"github.com/a1exCross/auth/internal/closer"
	"github.com/a1exCross/auth/internal/config"
	"github.com/a1exCross/auth/internal/repository"
	logsRepository "github.com/a1exCross/auth/internal/repository/logs"
	userRepository "github.com/a1exCross/auth/internal/repository/user"
	"github.com/a1exCross/auth/internal/service"
	userService "github.com/a1exCross/auth/internal/service/user"

	"context"
	"log"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient  db.Client
	txManager db.TxManager

	userRepo repository.UserRepository
	logsRepo repository.LogsRepository

	userServ service.UserService

	userImpl *userAPI.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %v", err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get gRPC config: %v", err)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %v", err)
		}

		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepo == nil {
		s.userRepo = userRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userRepo
}

func (s *serviceProvider) LogsRepository(ctx context.Context) repository.LogsRepository {
	if s.logsRepo == nil {
		s.logsRepo = logsRepository.NewRepository(s.DBClient(ctx))
	}

	return s.logsRepo
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userServ == nil {
		s.userServ = userService.NewService(
			s.UserRepository(ctx),
			s.TxManager(ctx),
			s.LogsRepository(ctx),
		)
	}

	return s.userServ
}

func (s *serviceProvider) UserImplementation(ctx context.Context) *userAPI.Implementation {
	if s.userImpl == nil {
		s.userImpl = userAPI.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}
