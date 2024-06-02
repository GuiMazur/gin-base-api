package dto

import "gin-base-api/pkg/models"

// Remember to add new validations tags to the validation message generator
type UpdateUserDto struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (updateUserDto *UpdateUserDto) ToUser() *models.User {
	return &models.User{
		Name:  updateUserDto.Name,
		Email: updateUserDto.Email,
	}
}
