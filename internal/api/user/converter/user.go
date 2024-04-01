package converter

import (
	"github.com/a1exCross/auth/internal/model"
	userPb "github.com/a1exCross/auth/pkg/user_v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserToProto - конвертирует модель пользователя в proto
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

// UserInfoToProto - конвертирует информацию о пользователе в proto
func UserInfoToProto(info model.UserInfo) *userPb.UserInfo {
	return &userPb.UserInfo{
		Username: info.Username,
		Name:     info.Name,
		Role:     userPb.UserRole(info.Role),
		Email:    info.Email,
	}
}

// ProtoToUserInfo - конвертирует proto в модель информации о пользователе
func ProtoToUserInfo(info *userPb.UserInfo) model.UserInfo {
	return model.UserInfo{
		Username: info.Username,
		Name:     info.Name,
		Role:     model.UserRole(info.Role),
		Email:    info.Email,
	}
}

// ProtoToUserInfoUpdate - конвертирует proto в модель информации о пользователе
func ProtoToUserInfoUpdate(info *userPb.UpdateInfo) model.UserInfo {
	return model.UserInfo{
		Name:  info.Name.Value,
		Role:  model.UserRole(info.Role),
		Email: info.Email.Value,
	}
}
