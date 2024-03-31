package app

import (
	"context"
	"log"

	accessAPI "github.com/a1exCross/auth/internal/api/access"
	authAPI "github.com/a1exCross/auth/internal/api/auth"
	userAPI "github.com/a1exCross/auth/internal/api/user"
	"github.com/a1exCross/auth/internal/config"
	"github.com/a1exCross/auth/internal/repository"
	logsRepository "github.com/a1exCross/auth/internal/repository/logs"
	userRepository "github.com/a1exCross/auth/internal/repository/user"
	"github.com/a1exCross/auth/internal/service"
	accessservice "github.com/a1exCross/auth/internal/service/access"
	authservice "github.com/a1exCross/auth/internal/service/auth"
	userService "github.com/a1exCross/auth/internal/service/user"
	"github.com/a1exCross/auth/internal/utils"

	"github.com/a1exCross/common/pkg/client/db"
	"github.com/a1exCross/common/pkg/client/db/pg"
	"github.com/a1exCross/common/pkg/client/db/transaction"
	"github.com/a1exCross/common/pkg/closer"
	"github.com/a1exCross/common/pkg/storage"
	cache "github.com/a1exCross/common/pkg/storage/redis"

	"github.com/go-redis/redis"
)

type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig
	redisConfig   config.RedisConfig
	jwtConfig     config.JWTConfig

	redisClient storage.Redis
	dbClient    db.Client
	txManager   db.TxManager

	userRepo repository.UserRepository
	logsRepo repository.LogsRepository

	userServ   service.UserService
	authServ   service.AuthService
	accessServ service.AccessService

	userImpl   *userAPI.Implementation
	authImpl   *authAPI.Implementation
	accessImpl *accessAPI.Implementation

	accessChecker utils.AccessChecker
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

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %v", err)
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := config.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %v", err)
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) RedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := config.NewRedisConfig()
		if err != nil {
			log.Fatalf("failed to get redis config: %v", err)
		}

		s.redisConfig = cfg
	}

	return s.redisConfig
}

func (s *serviceProvider) JWTConfig() config.JWTConfig {
	if s.jwtConfig == nil {
		cfg, err := config.NewJWTConfig()
		if err != nil {
			log.Fatalf("failed to get jwt config: %v", err)
		}

		s.jwtConfig = cfg
	}

	return s.jwtConfig
}

func (s *serviceProvider) RedisClient() storage.Redis {
	if s.redisClient == nil {
		cl, err := cache.NewRedisConnection(&redis.Options{
			Addr:     s.RedisConfig().Address(),
			Password: "",
			DB:       0,
		})
		if err != nil {
			log.Fatalf("failed to create redis client: %v", err)
		}

		err = cl.Ping()
		if err != nil {
			log.Fatalf("ping error: %v", err)
		}

		closer.Add(cl.Close)

		s.redisClient = cl
	}

	return s.redisClient
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
			s.RedisClient(),
		)
	}

	return s.userServ
}

func (s *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessServ == nil {
		s.accessServ = accessservice.NewService(
			s.JWTConfig(),
			s.AccessChecker(ctx))
	}

	return s.accessServ
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authServ == nil {
		s.authServ = authservice.NewService(
			s.RedisClient(),
			s.UserRepository(ctx),
			s.JWTConfig(),
		)
	}

	return s.authServ
}

func (s *serviceProvider) UserImplementation(ctx context.Context) *userAPI.Implementation {
	if s.userImpl == nil {
		s.userImpl = userAPI.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}

func (s *serviceProvider) AuthImplementation(ctx context.Context) *authAPI.Implementation {
	if s.authImpl == nil {
		s.authImpl = authAPI.NewImplementation(
			s.AuthService(ctx),
			s.AccessChecker(ctx),
		)
	}

	return s.authImpl
}

func (s *serviceProvider) AccessImplementation(ctx context.Context) *accessAPI.Implementation {
	if s.accessImpl == nil {
		s.accessImpl = accessAPI.NewImplementation(s.AccessService(ctx))
	}

	return s.accessImpl
}

func (s *serviceProvider) AccessChecker(ctx context.Context) utils.AccessChecker {
	if s.accessChecker == nil {
		s.accessChecker = utils.NewRouteAccessChecker(
			s.JWTConfig(),
			s.RedisClient(),
			s.UserRepository(ctx),
		)
	}

	return s.accessChecker
}
