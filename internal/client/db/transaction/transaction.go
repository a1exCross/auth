package transaction

import (
	"github.com/a1exCross/auth/internal/client/db"
	"github.com/a1exCross/auth/internal/client/db/pg"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"context"
)

type manager struct {
	db db.Transactor
}

// NewTransactionManager - создает новый менеджер транзакций
func NewTransactionManager(db db.Transactor) db.TxManager {
	return &manager{
		db: db,
	}
}

func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, fn db.Handler) (err error) {
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return errors.Wrap(err, "can`t begin transaction")
	}

	ctx = pg.MakeContext(ctx, tx)

	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("panic recovered: %v", err)
		}

		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Wrapf(err, "errRollback: %v", err)
			}

			return
		}

		if err == nil {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Wrap(err, "tx commit failed")
			}
		}
	}()

	if err = fn(ctx); err != nil {
		err = errors.Wrap(err, "failed executing code inside transaction")
	}

	return err
}

func (m *manager) ReadCommitted(ctx context.Context, fn db.Handler) error {
	txOpts := pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}
	return m.transaction(ctx, txOpts, fn)
}
