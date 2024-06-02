package exception

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Handle(ctx *gin.Context, err error) {
	if exception, ok := err.(*Exception); ok {
		ctx.JSON(exception.code, gin.H{"message": exception.message, "error": exception.error})
		return
	}
	
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Registro não encontrado.", "error": err.Error()})
		return
	}

	println("Untreated error: ", err.Error())
	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Erro interno.", "error": err.Error()})
}
