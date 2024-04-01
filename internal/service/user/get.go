package userservice

import (
	"context"
	"strconv"

	"github.com/a1exCross/auth/internal/model"

	"github.com/a1exCross/common/pkg/filter"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	var user *model.User

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		conditions := filter.MakeFilter(filter.Condition{
			Key:   model.IDFieldCode,
			Value: id,
		})

		user, errTx = s.userRepo.Get(ctx, conditions)
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
