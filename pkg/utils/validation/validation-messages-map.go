package validation

import "github.com/go-playground/validator/v10"

// Adicionar mensagens de validação customizadas para cada validator usado
var ValidationMessagesMap = map[string]func(validator.FieldError) string{
	"required": func(fieldError validator.FieldError) string {
		return "O campo " + fieldError.StructField() + " é obrigatório"
	},
	"datetime": func(fieldError validator.FieldError) string {
		return "O campo " + fieldError.StructField() + " deve seguir o padrão de data " + fieldError.Param()
	},
	"oneof": func(fieldError validator.FieldError) string {
		return "O campo " + fieldError.StructField() + " deve ser um dos seguintes valores: " + fieldError.Param()
	},
	"email": func(fieldError validator.FieldError) string {
		return "O campo " + fieldError.StructField() + " deve ser um e-mail válido"
	},
}
