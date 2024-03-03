package userservice

import (
	"github.com/a1exCross/auth/internal/model"

	"context"
	"strconv"
)

func (s *serv) Update(ctx context.Context, user *model.UserUpdate) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		errTx = s.userRepo.Update(ctx, user)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.logsRepo.Create(ctx, model.Log{
			Action:  "user updated",
			Content: strconv.FormatInt(user.ID, 10),
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
