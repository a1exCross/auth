package converter

import (
	"github.com/a1exCross/auth/internal/model"
	userPb "github.com/a1exCross/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserToProto(user *model.User) *userPb.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &userPb.User{
		Id:        user.ID,
		Info:      UserInfoToProto(user.Info),
		UpdatedAt: updatedAt,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}

func UserInfoToProto(info model.UserInfo) *userPb.UserInfo {
	return &userPb.UserInfo{
		Name:  info.Name,
		Role:  userPb.UserRole(info.Role),
		Email: info.Email,
	}
}
