package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorizedController struct {
	db map[string]string
}

func NewAuthorizedController(db map[string]string) ControllerInterface {
	return &AuthorizedController{
		db: db,
	}
}

func (this *AuthorizedController) Setup(router *gin.RouterGroup) {
	router.POST("admin", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})
}