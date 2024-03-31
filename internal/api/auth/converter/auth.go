package converter

import (
	"github.com/a1exCross/auth/internal/model"
	authpb "github.com/a1exCross/auth/pkg/auth_v1"
)

// AuthProtoToAuthDTO - конвертирует Proto в Model дл ручки Login
func AuthProtoToAuthDTO(req *authpb.LoginRequest) model.LoginDTO {
	return model.LoginDTO{
		Password: req.GetPassword(),
		Username: req.GetUsername(),
	}
}
