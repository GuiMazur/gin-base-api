package router

import (
	controllergroup "gin-base-api/pkg/controllers/controller-group"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(server *gin.Engine, db *gorm.DB) {
	controllerGroup := controllergroup.New(db)

	controllerGroup.AppController.Setup(server.Group("/"))

	controllerGroup.UserController.Setup(server.Group("/user"))
}
