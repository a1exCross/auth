package tests

import (
	"context"
	"errors"
	"testing"

	userapi "github.com/a1exCross/auth/internal/api/user"
	"github.com/a1exCross/auth/internal/model"
	"github.com/a1exCross/auth/internal/service"
	"github.com/a1exCross/auth/internal/service/mocks"
	"github.com/a1exCross/auth/pkg/user_v1"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)
	ctx := context.Background()

	type mockAction func(mc minimock.MockController) service.UserService

	correctReq := user_v1.CreateRequest{
		Info: &user_v1.UserInfo{
			Role:  user_v1.UserRole(1),
			Email: "email",
			Name:  "name",
		},
		Pass: &user_v1.UserPassword{
			Password:        "password",
			PasswordConfirm: "password",
		},
	}

	incorrectReq := user_v1.CreateRequest{
		Info: &user_v1.UserInfo{
			Role:  user_v1.UserRole(1),
			Email: "email",
			Name:  "name",
		},
		Pass: &user_v1.UserPassword{
			Password:        "password",
			PasswordConfirm: "password12345",
		},
	}

	paramsCreate := model.UserCreate{
		Info: model.UserInfo{
			Role:  model.UserRole(1),
			Email: "email",
			Name:  "name",
		},
		Password: "password",
	}

	id := int64(1)

	resp := user_v1.CreateResponse{
		Id: id,
	}

	tests := []struct {
		name       string
		ctx        context.Context
		req        *user_v1.CreateRequest
		err        error
		expected   *user_v1.CreateResponse
		mockAction mockAction
	}{
		{
			name:     "sucessfull test",
			req:      &correctReq,
			err:      nil,
			ctx:      ctx,
			expected: &resp,
			mockAction: func(mc minimock.MockController) service.UserService {
				userServiceMock := mocks.NewUserServiceMock(mc)
				userServiceMock.CreateMock.Expect(ctx, &paramsCreate).Return(id, nil)

				return userServiceMock
			},
		},
		{
			name:     "mismatch passwords",
			req:      &incorrectReq,
			err:      errors.New("passwords mismatch"),
			ctx:      ctx,
			expected: nil,
			mockAction: func(mc minimock.MockController) service.UserService {
				userServiceMock := mocks.NewUserServiceMock(mc)

				return userServiceMock
			},
		},
		{
			name:     "some error",
			req:      &correctReq,
			err:      errors.New("failed to create user: error"),
			ctx:      ctx,
			expected: nil,
			mockAction: func(mc minimock.MockController) service.UserService {
				userServiceMock := mocks.NewUserServiceMock(mc)
				userServiceMock.CreateMock.Expect(ctx, &paramsCreate).Return(0, errors.New("error"))

				return userServiceMock
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			userServ := test.mockAction(mc)
			impl := userapi.NewImplementation(userServ)

			res, err := impl.Create(test.ctx, test.req)

			require.Equal(t, res, test.expected)
			if err != nil && test.err != nil {
				require.Equal(t, test.err.Error(), err.Error())
			} else {
				require.Equal(t, test.err, err)
			}
		})
	}
}
