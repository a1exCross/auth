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
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)
	ctx := context.Background()

	type mockAction func(mc minimock.MockController) service.UserService
	type mockAccess func(mc minimock.MockController) service.AccessService

	id := int64(1)

	correctReq := &user_v1.UpdateRequest{
		Id: id,
		Info: &user_v1.UpdateInfo{
			Name: &wrappers.StringValue{
				Value: "name",
			},
			Email: &wrappers.StringValue{
				Value: "email",
			},
			Role: user_v1.UserRole(1),
		},
	}

	updateDTOData := &model.UserUpdate{
		Info: model.UserInfo{
			Role:  model.UserRole(1),
			Email: "email",
			Name:  "name",
		},
		ID: id,
	}

	tests := []struct {
		name       string
		ctx        context.Context
		req        *user_v1.UpdateRequest
		err        error
		expected   *empty.Empty
		mockAction mockAction
		mockAccess
	}{
		{
			name:     "sucessfull test",
			req:      correctReq,
			err:      nil,
			ctx:      ctx,
			expected: &empty.Empty{},
			mockAction: func(mc minimock.MockController) service.UserService {
				userServiceMock := mocks.NewUserServiceMock(mc)
				userServiceMock.UpdateMock.Expect(ctx, updateDTOData).Return(nil)

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
			err:      errors.New("failed to update user: error"),
			ctx:      ctx,
			expected: nil,
			mockAction: func(mc minimock.MockController) service.UserService {
				userServiceMock := mocks.NewUserServiceMock(mc)
				userServiceMock.UpdateMock.Expect(ctx, updateDTOData).Return(errors.New("error"))

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

			res, err := impl.Update(test.ctx, test.req)

			require.Equal(t, res, test.expected)
			if err != nil && test.err != nil {
				require.Equal(t, test.err.Error(), err.Error())
			} else {
				require.Equal(t, test.err, err)
			}
		})
	}
}
