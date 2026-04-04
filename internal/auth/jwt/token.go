package jwt

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/domain"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/repository/models"
)

func GenerateToken(u *models.GetUserByEmailOut, cfg *config.JWT) (string, error) {
	expAt := time.Now().Add(time.Second * time.Duration(cfg.ExpirationTime)).Unix() // tiempo de expiración del token

	secretKey := cfg.SecretKey // clave secreta para firmar el token

	iat := time.Now().Unix() // hora en que se emitió el token

	claims := Claims{
		UserID: u.ID,
		Email:  u.Email,
		Role:   u.RoleName,
		Status: string(u.Status),
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  iat,
			ExpiresAt: expAt,
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := newToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("could not sign JWT token: %w", err)
	}

	return token, nil
}

func validateToken(r *http.Request, cfg *config.JWT) (domain.User, error) {
	token, err := getToken(r)
	if err != nil {
		log.Printf("Error retrieving token: %v", err)
		return domain.User{}, fmt.Errorf("no token found in request: %w", err)
	}

	// jwt.Parse recibe un keyfunc con firma func(*jwt.Token) (any, error)
	// usamos un closure para capturar cfg sin cambiar la firma esperada
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return validateMethodAndGetSecret(t, cfg)
	})
	if err != nil {
		log.Printf("Token not valid: %v\n", err)
		return domain.User{}, err
	}

	userData, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		log.Println("Unable to retrieve payload information or token is invalid")
		return domain.User{}, fmt.Errorf("invalid token claims")
	}

	_, ok = userData["email"].(string)
	if !ok {
		log.Println("Email field missing or not a string in token claims")
		return domain.User{}, fmt.Errorf("email field is missing or invalid in token claims")
	}

	return domain.User{
		Email: userData["email"].(string),
	}, nil
}

func getToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header missing")
	}

	parts := strings.SplitN(authHeader, " ", 2)

	// Validamos que existan dos partes y que la primera sea "bearer" (sin importar mayúsculas)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
		return "", fmt.Errorf("invalid authorization header format")
	}

	token := parts[1]

	return token, nil
}

func validateMethodAndGetSecret(token *jwt.Token, cfg *config.JWT) (any, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(cfg.SecretKey), nil
}
