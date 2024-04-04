package userservice

import (
	"github.com/a1exCross/auth/internal/repository"
	"github.com/a1exCross/auth/internal/service"

	"github.com/a1exCross/common/pkg/client/db"
	"github.com/a1exCross/common/pkg/storage"
)

type serv struct {
	userRepo  repository.UserRepository
	logsRepo  repository.LogsRepository
	txManager db.TxManager
	storage   storage.Redis
}

// NewService - создает сервисный слой для работы с пользователями
func NewService(userRepo repository.UserRepository, tx db.TxManager, logRepo repository.LogsRepository, storage storage.Redis) service.UserService {
	return &serv{
		userRepo:  userRepo,
		txManager: tx,
		logsRepo:  logRepo,
		storage:   storage,
	}
}
