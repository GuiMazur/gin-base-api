package dtos

import "gin-base-api/models"

type CreateUserDto struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age   uint8  `json:"age"`
}

func (this *CreateUserDto) ToUser() *models.User {
	return &models.User{
		Name:  this.Name,
		Email: this.Email,
		Age:   &this.Age,
	}
}