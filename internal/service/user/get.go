package userservice

import (
	"github.com/a1exCross/auth/internal/model"

	"context"
	"strconv"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	var user *model.User

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		user, errTx = s.userRepo.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.logsRepo.Create(ctx, model.Log{
			Action:  "user fetch",
			Content: strconv.FormatInt(id, 10),
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
