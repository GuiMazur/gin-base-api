package main

import (
	"fmt"
	"gin-base-api/config"
	"gin-base-api/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	config := config.NewConfig()

	server := gin.Default()

	server.Use(Middleware1())
	
	router.SetupRouter(server)
	
	server.Run(config.App.Host + ":" + config.App.Port)
}

func Middleware1() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		fmt.Println("Middleware 1")
	}
}