package router

import (
	"gin-base-api/controllers"
	"gin-base-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(server *gin.Engine, db *gorm.DB) {
	var (
		appController = controllers.NewAppController(services.NewAppService(db))
		authorizedController = controllers.NewAuthorizedController(db)
	)

	appController.Setup(server.Group("/"))

	authorizedGroup := server.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorizedController.Setup(authorizedGroup)
}