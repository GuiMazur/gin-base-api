package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Unauthorize(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Sem permissão."})
}
