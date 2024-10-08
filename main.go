package main

import (
	"gin-base-api/pkg/config"
	"gin-base-api/pkg/db"
	"gin-base-api/pkg/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	config := config.New()

	db.New()

	server := gin.Default()

	router.Setup(server)

	server.Run(config.App.Host + ":" + config.App.Port)
}
