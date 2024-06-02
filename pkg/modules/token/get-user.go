package token

import (
	"gin-base-api/pkg/utils/exception"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) (*UserClaims, error) {
	var user *UserClaims

	userData, ok := ctx.Get("user")
	if !ok {
		return nil, exception.New("Sem permissão.", 401, "")
	}

	user, ok = userData.(*UserClaims)
	if !ok {
		return nil, exception.New("Sem permissão.", 401, "")
	}

	return user, nil
}
