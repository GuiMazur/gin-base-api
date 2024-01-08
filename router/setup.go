package router

import (
	"gin-base-api/controllers"
	"gin-base-api/services"

	"github.com/gin-gonic/gin"
)

var (
	db = make(map[string]string)
	appController = controllers.NewAppController(services.NewAppService(db))
	authorizedController = controllers.NewAuthorizedController(db)
)

func SetupRouter(server *gin.Engine) {
	appController.Setup(server.Group("/"))

	authorizedGroup := server.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorizedController.Setup(authorizedGroup)
}