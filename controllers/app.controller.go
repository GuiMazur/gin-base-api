package controllers

import (
 	"gin-base-api/controllers/dtos/app"
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
		name := ctx.Params.ByName("name")
		ret, err := this.AppService.GetUser(name)
		if err != nil {
			exception.HandleException(ctx, err)
			return
		}
		ctx.JSON(200, ret)
	})

	router.POST("/user", func(ctx *gin.Context) {
		var createUserDto dtos.CreateUserDto;
		if err := ctx.ShouldBindJSON(&createUserDto); err != nil {
			exception.HandleException(ctx, exception.NewException(err.Error(), 400))
			return
		}
		ret, err := this.AppService.CreateUser(&createUserDto)
		if err != nil {
			exception.HandleException(ctx, err)
			return
		}
		ctx.JSON(200, ret)
	})
}
