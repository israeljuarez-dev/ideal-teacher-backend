package validator

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	v *validator.Validate
}

func New() *Validator {
	return &Validator{
		v: validator.New(),
	}
}

// Validate valida una struct y retorna un slice de FieldError si hay errores
func (val *Validator) Validate(a any) []FieldError {
	err := val.v.Struct(a)
	if err == nil {
		return nil
	}

	var ve validator.ValidationErrors

	if !errors.As(err, &ve) {
		return []FieldError{
			{
				Field:   "unknown",
				Message: "invalid input",
			},
		}
	}

	out := make([]FieldError, 0, len(ve))

	for _, fe := range ve {
		out = append(out, FieldError{
			Field:   toSnakeCase(fe.Field()),
			Message: fieldMessage(fe),
		})
	}

	return out
}

func fieldMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Este campo es requerido"
	case "email":
		return "Debe ser un correom electrónico válido"
	case "min":
		return "Debe tener al menos" + fe.Param() + " carácteres"
	case "max":
		return "No debe exceder " + fe.Param() + " carácteres"
	default:
		return "Inválido"
	}
}

// toSnakeCase convierte "FirstName" → "first_name"
func toSnakeCase(s string) string {
	out := make([]rune, 0, len(s)+4)
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				out = append(out, '_')
			}
			out = append(out, r+32)
		} else {
			out = append(out, r)
		}
	}
	return string(out)
}
