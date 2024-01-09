package services

import (
	"gin-base-api/controllers/dtos/app"
	"gin-base-api/exception"
	"gin-base-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppService struct {
	db *gorm.DB
}

func NewAppService(db *gorm.DB) *AppService {
	return &AppService{
		db: db,
	}
}

func (this *AppService) Ping() (string, error) {
	return "pong", nil
}

func (this *AppService) GetUser(name string) (gin.H, error) {
	var user models.User
	result := this.db.Where(&models.User{Name: name}).First(&user)
	if result.Error != nil {
		return nil, exception.NewException("User not found", 404)
	}
	return gin.H{"user": user}, nil
}

func (this *AppService) CreateUser(createUserDto *dtos.CreateUserDto) (gin.H, error) {
	user := createUserDto.ToUser()
	result := this.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return gin.H{"user": user}, nil
}