package userservice

import (
	"github.com/a1exCross/auth/internal/client/db"
	"github.com/a1exCross/auth/internal/repository"
	"github.com/a1exCross/auth/internal/service"
)

type serv struct {
	userRepo  repository.UserRepository
	logsRepo  repository.LogsRepository
	txManager db.TxManager
}

// NewService - создает сервисный слой для работы с пользователями
func NewService(userRepo repository.UserRepository, tx db.TxManager, logRepo repository.LogsRepository) service.UserService {
	return &serv{
		userRepo:  userRepo,
		txManager: tx,
		logsRepo:  logRepo,
	}
}
