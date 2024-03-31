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
	"github.com/a1exCross/common/pkg/storage"
	storagemocks "github.com/a1exCross/common/pkg/storage/mocks"

	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	type dbClientMock func(mc *minimock.Controller) db.Client
	type txManagerMock func(mc *minimock.Controller) db.TxManager
	type storageMock func(mc *minimock.Controller) storage.Redis

	ctx := context.Background()
	mc := minimock.NewController(t)
	id := int64(1)

	//timeNow := time.Now()

	userDTO := &model.UserCreate{
		Info: model.UserInfo{
			Email: "email",
			Role:  model.UserRole(1),
			Name:  "name",
		},
		Password: "pass",
	}

	/*	hash, _ := utils.HashPassword("pass")

		user := &model.User{
			Info: model.UserInfo{
				Email: "email",
				Role:  model.UserRole(1),
				Name:  "name",
			},
			Password:  hash,
			CreatedAt: timeNow,
			UpdatedAt: sql.NullTime{},
			ID:        id,
		}*/

	tests := []struct {
		name      string
		err       error
		expected  int64
		dbClient  dbClientMock
		txManager txManagerMock
		storageMock
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
					return pgx.ErrNoRows
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
			storageMock: func(mc *minimock.Controller) storage.Redis {
				mock := storagemocks.NewRedisMock(mc)

				return mock
			},
			err:      nil,
			expected: id,
		},
		{
			name: "error at create",
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

					return errors.New("failed to scan")
				})

				dbb.QueryRowContextMock.Return(row)
				dbb.ScanOneContextMock.Set(func(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) (err error) {
					return pgx.ErrNoRows
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

				tx.RollbackMock.Return(nil)

				transactor.BeginTxMock.Expect(ctx, txOptions).Return(tx, nil)

				txManager := transaction.NewTransactionManager(transactor)

				return txManager
			},
			storageMock: func(mc *minimock.Controller) storage.Redis {
				mock := storagemocks.NewRedisMock(mc)

				return mock
			},
			err:      errors.New("failed executing code inside transaction: error at query to database: failed to scan"),
			expected: 0,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			dbMockClient := test.dbClient(mc)
			txManager := test.txManager(mc)
			redisMock := test.storageMock(mc)

			userRepo := userRepository.NewRepository(dbMockClient)
			logRepo := logsRepository.NewRepository(dbMockClient)

			userServ := userservice.NewService(userRepo, txManager, logRepo, redisMock)

			res, err := userServ.Create(ctx, userDTO)

			require.Equal(t, test.expected, res)

			if err != nil && test.err != nil {
				require.Equal(t, test.err.Error(), err.Error())
			} else {
				require.Equal(t, test.err, err)
			}
		})
	}
}
