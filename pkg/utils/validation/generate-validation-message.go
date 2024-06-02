package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func GenerateValidationMessage(validationErrors validator.ValidationErrors) string {
	validationMessages := make([]string, len(validationErrors))

	for i, validationError := range validationErrors {
		messageGenerator := ValidationMessagesMap[validationError.Tag()]

		if messageGenerator != nil {
			validationMessages[i] = messageGenerator(validationError)
		} else {
			validationMessages[i] = "Erro com o campo " + validationError.StructField() + ". Tag: " + validationError.Tag()
		}
	}

	validationMessage := "Erro de validação: " + strings.Join(validationMessages, " | ") + "."

	return validationMessage
}
