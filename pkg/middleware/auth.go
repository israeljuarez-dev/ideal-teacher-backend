package middleware

import (
	"log/slog"
	"net/http"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/jwt"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
	"github.com/israeljuarez-dev/ideal-teacher-backend/pkg/response"
)

func ValidateJWT(cfg *config.JWT, log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := jwt.ValidateToken(r, log, cfg)
			if err != nil {
				log.Warn("ValidateJWT: invalid token", "error", err, "path", r.URL.Path)
				response.Error(w, response.ErrorResponse{
					Status:  http.StatusUnauthorized,
					Message: "invalid or missing token",
				})
				return
			}

			log.Debug("ValidateJWT: token valid", "path", r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}
}
