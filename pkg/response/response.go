package response

import (
	"encoding/json"
	"net/http"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/validator"
)

// JSON escribe una respuesta exitosa con datos genéricos
func JSON[T any](w http.ResponseWriter, s Success[T]) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s.Status)
	_ = json.NewEncoder(w).Encode(Success[T]{
		Status:  s.Status,
		Message: s.Message,
		Data:    s.Data,
	})
}

// Error escribe una respuesta de error simple (sin campos)
func Error(w http.ResponseWriter, e ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Status:  e.Status,
		Message: e.Message,
	})
}

// ValidationError escribe errores de validación estructurados
func ValidationError(w http.ResponseWriter, fields []validator.FieldError) {
	details := make([]ErrorDetail, 0, len(fields))
	for _, f := range fields {
		details = append(details, ErrorDetail{
			Field:   f.Field,
			Message: f.Message,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Status:  http.StatusUnprocessableEntity,
		Message: "validation failed",
		Errors:  details,
	})
}
