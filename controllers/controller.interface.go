package controllers

import "github.com/gin-gonic/gin"

type ControllerInterface interface {
	Setup(router *gin.RouterGroup)
}

