package tests

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/a1exCross/auth/internal/model"
	logsRepository "github.com/a1exCross/auth/internal/repository/logs"
	userRepository "github.com/a1exCross/auth/internal/repository/user"
	userservice "github.com/a1exCross/auth/internal/service/user"

	"github.com/a1exCross/common/pkg/client/db"
	dbmocks "github.com/a1exCross/common/pkg/client/db/mocks"
	"github.com/a1exCross/common/pkg/client/db/transaction"

	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	type dbClientMock func(mc *minimock.Controller) db.Client
	type txManagerMock func(mc *minimock.Controller) db.TxManager

	ctx := context.Background()
	mc := minimock.NewController(t)
	id := int64(1)
	timeNow := time.Now()

	user := &model.User{
		ID: id,
		Info: model.UserInfo{
			Name:  "name",
			Email: "email",
			Role:  0,
		},
		UpdatedAt: sql.NullTime{
			Valid: true,
			Time:  timeNow,
		},
		CreatedAt: timeNow,
	}

	tests := []struct {
		name      string
		err       error
		expected  *model.User
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
				dbb.ScanOneContextMock.Set(func(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) (err error) {
					res, ok := dest.(*model.User)
					if ok {
						res.Info = user.Info
						res.CreatedAt = user.CreatedAt
						res.UpdatedAt = user.UpdatedAt
					}

					_ = res

					return nil
				})

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
			expected: &model.User{
				ID: id,
				Info: model.UserInfo{
					Name:  "name",
					Email: "email",
					Role:  0,
				},
				UpdatedAt: sql.NullTime{
					Valid: true,
					Time:  timeNow,
				},
				CreatedAt: timeNow,
			},
		},
		{
			name: "error at get",
			dbClient: func(mc *minimock.Controller) db.Client {
				client := dbmocks.NewClientMock(mc)

				return client
			},
			txManager: func(mc *minimock.Controller) db.TxManager {

				transactor := dbmocks.NewTransactorMock(mc)

				txOptions := pgx.TxOptions{
					IsoLevel: pgx.ReadCommitted,
				}

				transactor.BeginTxMock.Expect(ctx, txOptions).Return(nil, errors.New("tx error"))

				txManager := transaction.NewTransactionManager(transactor)

				return txManager
			},
			err:      errors.New("can`t begin transaction: tx error"),
			expected: nil,
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

			res, err := userServ.Get(ctx, id)

			require.Equal(t, test.expected, res)

			if err != nil && test.err != nil {
				require.Equal(t, test.err.Error(), err.Error())
			} else {
				require.Equal(t, test.err, err)
			}
		})
	}
}
