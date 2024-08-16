package router

import (
	"gin-base-api/pkg/modules/app"
	"gin-base-api/pkg/modules/user"

	"github.com/gin-gonic/gin"
)

func Setup(server *gin.Engine) {
	app.NewController().Setup(server.Group("/"))

	user.NewController().Setup(server.Group("/user"))
}
