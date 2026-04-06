package response

// ErrorDetail describe un error individual
type ErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

// Error es la envoltura para respuestas de error
type ErrorResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Errors  []ErrorDetail `json:"errors,omitempty"`
}
