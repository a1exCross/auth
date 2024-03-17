package tests

import (
	"context"
	"testing"

	"github.com/a1exCross/auth/internal/model"
	logsRepository "github.com/a1exCross/auth/internal/repository/logs"
	userRepository "github.com/a1exCross/auth/internal/repository/user"
	userservice "github.com/a1exCross/auth/internal/service/user"

	"github.com/a1exCross/common/pkg/client/db"
	dbmocks "github.com/a1exCross/common/pkg/client/db/mocks"
	"github.com/a1exCross/common/pkg/client/db/transaction"

	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	type dbClientMock func(mc *minimock.Controller) db.Client
	type txManagerMock func(mc *minimock.Controller) db.TxManager

	ctx := context.Background()
	mc := minimock.NewController(t)
	id := int64(1)

	userDTO := &model.UserUpdate{
		Info: model.UserInfo{
			Email: "email",
			Role:  model.UserRole(1),
			Name:  "name",
		},
	}

	tests := []struct {
		name      string
		err       error
		expected  int64
		dbClient  dbClientMock
		txManager txManagerMock
	}{
		{
			name: "successfull test",
			dbClient: func(mc *minimock.Controller) db.Client {
				client := dbmocks.NewClientMock(mc)
				dbb := dbmocks.NewDBMock(mc)
				row := dbmocks.NewRowMock(mc)

				row.ScanMock.Set(func(dest ...interface{}) (err error) {
					res, ok := dest[0].(*int64)
					if ok {
						*res = id
					}

					_ = res

					return nil
				})

				dbb.QueryRowContextMock.Return(row)
				dbb.ExecContextMock.Return(pgconn.CommandTag{}, nil)

				client.DBMock.Return(dbb)

				return client
			},
			txManager: func(mc *minimock.Controller) db.TxManager {
				tx := dbmocks.NewTxMock(mc)
				transactor := dbmocks.NewTransactorMock(mc)

				txOptions := pgx.TxOptions{
					IsoLevel: pgx.ReadCommitted,
				}

				tx.CommitMock.Return(nil)
				transactor.BeginTxMock.Expect(ctx, txOptions).Return(tx, nil)

				txManager := transaction.NewTransactionManager(transactor)

				return txManager
			},
			err: nil,
		},
		{
			name: "error at update",
			dbClient: func(mc *minimock.Controller) db.Client {
				client := dbmocks.NewClientMock(mc)
				dbb := dbmocks.NewDBMock(mc)
				row := dbmocks.NewRowMock(mc)

				row.ScanMock.Set(func(dest ...interface{}) (err error) {
					res, ok := dest[0].(*int64)
					if ok {
						*res = id
					}

					_ = res

					return nil
				})

				dbb.QueryRowContextMock.Return(row)
				dbb.ExecContextMock.Return(pgconn.CommandTag{}, nil)

				client.DBMock.Return(dbb)

				return client
			},
			txManager: func(mc *minimock.Controller) db.TxManager {
				tx := dbmocks.NewTxMock(mc)
				transactor := dbmocks.NewTransactorMock(mc)

				txOptions := pgx.TxOptions{
					IsoLevel: pgx.ReadCommitted,
				}

				tx.CommitMock.Return(errors.New("commit error"))

				transactor.BeginTxMock.Expect(ctx, txOptions).Return(tx, nil)

				txManager := transaction.NewTransactionManager(transactor)

				return txManager
			},
			err: errors.New("tx commit failed: commit error"),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			dbMockClient := test.dbClient(mc)
			txManager := test.txManager(mc)

			userRepo := userRepository.NewRepository(dbMockClient)
			logRepo := logsRepository.NewRepository(dbMockClient)

			userServ := userservice.NewService(userRepo, txManager, logRepo)

			err := userServ.Update(ctx, userDTO)

			if err != nil && test.err != nil {
				require.Equal(t, test.err.Error(), err.Error())
			} else {
				require.Equal(t, test.err, err)
			}
		})
	}
}
