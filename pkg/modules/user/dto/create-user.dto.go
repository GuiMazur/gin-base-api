package dto

import "gin-base-api/pkg/models"

// Remember to add new validations tags to the validation message generator
type CreateUserDto struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (createUserDto *CreateUserDto) ToUser() *models.User {
	return &models.User{
		Name:     createUserDto.Name,
		Email:    createUserDto.Email,
		Password: createUserDto.Password,
	}
}
