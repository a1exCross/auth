package converter

import (
	"github.com/a1exCross/auth/internal/model"
	accesspb "github.com/a1exCross/auth/pkg/access_v1"
)

// RolesFromProto - конвертирует роли пользователей в Proto
func RolesFromProto(roles []accesspb.UserRole) []model.UserRole {
	var converted []model.UserRole

	for _, role := range roles {
		converted = append(converted, model.UserRole(role.Number()))
	}

	return converted
}
