package middleware

import (
	"net/http"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
)

func ValidateJWT(cfg *config.JWT) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}
