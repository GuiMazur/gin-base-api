package middlewares

import (
	"fmt"
	"gin-base-api/pkg/modules/token"
	"gin-base-api/pkg/utils/exception"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	tokenService := token.NewService()

	return func(ctx *gin.Context) {
		accessToken, err := ctx.Cookie("accessToken")

		if err != nil {
			fmt.Println(err)
			exception.Unauthorize(ctx)
			ctx.Abort()
			return
		}

		userClaims, err := tokenService.ParseAccessToken(accessToken)

		if err != nil {
			fmt.Println(err)
			exception.Unauthorize(ctx)
			ctx.Abort()
			return
		}

		ctx.Set("user", userClaims)

		ctx.Next()
	}
}
