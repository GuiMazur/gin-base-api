package dto

import "gin-base-api/pkg/models"

// Remember to add new validations tags to the validation message generator
type UpdatePasswordDto struct {
	NewPassword string `json:"newPassword" validate:"required"`
	OldPassword string `json:"oldPassword" validate:"required"`
}

func (updatePasswordDto *UpdatePasswordDto) ToUser() *models.User {
	return &models.User{
		Password: updatePasswordDto.NewPassword,
	}
}
