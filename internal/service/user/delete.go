package userservice

import (
	"github.com/a1exCross/auth/internal/model"

	"context"
	"strconv"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		errTx = s.userRepo.Delete(ctx, id)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.logsRepo.Create(ctx, model.Log{
			Action:  "user deleted",
			Content: strconv.FormatInt(id, 10),
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
