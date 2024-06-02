package app

import (
	"gin-base-api/pkg/utils/exception"

	"github.com/gin-gonic/gin"
)

type AppController struct {
	*AppService
}

var appControllerInstance *AppController

func NewController(appService *AppService) *AppController {
	if appControllerInstance == nil {
		appControllerInstance = &AppController{
			AppService: appService,
		}
	}
	return appControllerInstance
}

func (appController *AppController) Setup(router *gin.RouterGroup) {
	router.GET("", func(ctx *gin.Context) {
		ret, err := appController.AppService.Ping()
		if err != nil {
			exception.Handle(ctx, err)
			return
		}
		ctx.JSON(200, ret)
	})
}
