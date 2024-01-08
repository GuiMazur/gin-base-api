package services

import (
	"gin-base-api/exception"

	"github.com/gin-gonic/gin"
)

type AppService struct {
	db map[string]string
}

func NewAppService(db map[string]string) *AppService {
	return &AppService{
		db: db,
	}
}

func (this *AppService) Ping() (string, error) {
	return "pong", nil
}

func (this *AppService) GetUser(user string) (gin.H, error) {
	value, ok := this.db[user]
	if !ok {
		return nil, exception.NewException("User not found", 404)
	}
	return gin.H{"user": user, "status": value}, nil
}