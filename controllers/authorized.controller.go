package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthorizedController struct {
	db *gorm.DB
}

func NewAuthorizedController(db *gorm.DB) ControllerInterface {
	return &AuthorizedController{
		db: db,
	}
}

func (this *AuthorizedController) Setup(router *gin.RouterGroup) {
	router.POST("admin", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})
}