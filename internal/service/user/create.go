package userservice

import (
	"context"
	"fmt"
	"strconv"

	"github.com/a1exCross/auth/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *serv) Create(ctx context.Context, userParams *model.UserCreate) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userParams.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("failed hash password: %v", err)
	}

	userParams.Password = string(hashedPassword)

	var id int64

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		id, errTx = s.userRepo.Create(ctx, userParams)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.logsRepo.Create(ctx, model.Log{
			Action:  "user created",
			Content: strconv.FormatInt(id, 10),
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}
