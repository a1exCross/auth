package tests

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	userapi "github.com/a1exCross/auth/internal/api/user"
	"github.com/a1exCross/auth/internal/model"
	"github.com/a1exCross/auth/internal/service"
	"github.com/a1exCross/auth/internal/service/mocks"
	"github.com/a1exCross/auth/pkg/user_v1"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)
	ctx := context.Background()

	type mockAction func(mc minimock.MockController) service.UserService
	type mockAccess func(mc minimock.MockController) service.AccessService

	correctReq := &user_v1.GetRequest{
		Id: 1,
	}

	id := int64(1)

	timeNow := time.Now()

	resGet := &model.User{
		ID: id,
		Info: model.UserInfo{
			Role:  model.UserRole(1),
			Email: "email",
			Name:  "name",
		},
		UpdatedAt: sql.NullTime{
			Time:  timeNow,
			Valid: true,
		},
		CreatedAt: timeNow,
	}

	resp := &user_v1.GetResponse{
		User: &user_v1.User{
			Id: id,
			Info: &user_v1.UserInfo{
				Role:  user_v1.UserRole(1),
				Email: "email",
				Name:  "name",
			},
			CreatedAt: timestamppb.New(timeNow),
			UpdatedAt: timestamppb.New(timeNow),
		},
	}

	tests := []struct {
		name       string
		ctx        context.Context
		req        *user_v1.GetRequest
		err        error
		expected   *user_v1.GetResponse
		mockAction mockAction
		mockAccess
	}{
		{
			name:     "sucessfull test",
			req:      correctReq,
			err:      nil,
			ctx:      ctx,
			expected: resp,
			mockAction: func(mc minimock.MockController) service.UserService {
				userServiceMock := mocks.NewUserServiceMock(mc)
				userServiceMock.GetMock.Expect(ctx, id).Return(resGet, nil)

				return userServiceMock
			},
			mockAccess: func(mc minimock.MockController) service.AccessService {
				mock := mocks.NewAccessServiceMock(mc)

				return mock
			},
		},
		{
			name:     "some error",
			req:      correctReq,
			err:      errors.New("failed to get user: error"),
			ctx:      ctx,
			expected: nil,
			mockAction: func(mc minimock.MockController) service.UserService {
				userServiceMock := mocks.NewUserServiceMock(mc)
				userServiceMock.GetMock.Expect(ctx, id).Return(nil, errors.New("error"))

				return userServiceMock
			},
			mockAccess: func(mc minimock.MockController) service.AccessService {
				mock := mocks.NewAccessServiceMock(mc)

				return mock
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			userServ := test.mockAction(mc)
			accessServ := test.mockAccess(mc)

			impl := userapi.NewImplementation(userServ, accessServ)

			res, err := impl.Get(test.ctx, test.req)

			require.Equal(t, res, test.expected)
			if err != nil && test.err != nil {
				require.Equal(t, test.err.Error(), err.Error())
			} else {
				require.Equal(t, test.err, err)
			}
		})
	}
}
