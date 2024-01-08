package controllers

import (
	"gin-base-api/exception"
	"gin-base-api/services"

	"github.com/gin-gonic/gin"
)

type AppController struct {
	*services.AppService
}

func NewAppController(appService *services.AppService) ControllerInterface {
	return &AppController{
		AppService: appService,
	}
}

func (this *AppController) Setup(router *gin.RouterGroup) {
	router.GET("/", func(ctx *gin.Context) {
		ret, err := this.AppService.Ping()
		if err != nil {
			exception.HandleException(ctx, err)
			return
		}
		ctx.JSON(200, ret)
	})

	router.GET("/user/:name", func(ctx *gin.Context) {
		user := ctx.Params.ByName("name")
		ret, err := this.AppService.GetUser(user)
		if err != nil {
			exception.HandleException(ctx, err)
			return
		}
		ctx.JSON(200, ret)
	})
}
