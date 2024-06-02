package dto

// Remember to add new validations tags to the validation message generator
type LoginDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
