package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Exception struct {
	message string
	code int
}

func NewException(message string, code int) *Exception {
	return &Exception{
		message: message,
		code: code,
	}
}

func (e *Exception) Error() string {
	return e.message
}

func HandleException(ctx *gin.Context, err error) {
	if exception, ok := err.(*Exception); ok {
		ctx.JSON(exception.code, gin.H{"error": exception.Error()})
		return
	}
	println("Untreated error: ", err.Error())
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
}