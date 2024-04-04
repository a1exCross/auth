package userservice

import (
	"context"
	"strconv"

	"github.com/a1exCross/auth/internal/model"
	"github.com/a1exCross/auth/internal/utils"

	"github.com/a1exCross/common/pkg/filter"

	"github.com/pkg/errors"
)

func (s *serv) Create(ctx context.Context, userParams *model.UserCreate) (int64, error) {
	conditions := filter.MakeFilter(filter.Condition{
		Key:   model.UserNameFieldCode,
		Value: userParams.Info.Username,
	})

	user, err := s.userRepo.Get(ctx, conditions)
	if err != nil && err.Error() != utils.UserNotFound {
		return 0, err
	}

	if user != nil {
		return 0, errors.Errorf(`user with username "%s" already exist`, userParams.Info.Username)
	}

	hashedPassword, err := utils.HashPassword(userParams.Password)
	if err != nil {
		return 0, errors.Errorf("failed hash password: %v", err)
	}

	userParams.Password = hashedPassword

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
