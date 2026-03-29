package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Env = env{}
)

func New() (env, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error cargando .env: %v", err)
		return env{}, err
	}

	Env = env{
		Container: Container{
			App: App{
				Name:           os.Getenv("APP_NAME"),
				Protocol:       os.Getenv("APP_PROTOCOL"),
				Port:           os.Getenv("SERVER_PORT"),
				AllowedOrigins: os.Getenv("ALLOWED_ORIGINS"),
				AllowedMethods: os.Getenv("ALLOWED_METHODS"),
			},
			DB: DB{
				PostgresEnv: PostgresEnv{
					Host:     os.Getenv("DB_HOST"),
					Port:     os.Getenv("DB_PORT"),
					Name:     os.Getenv("DB_NAME"),
					User:     os.Getenv("DB_USER"),
					Password: os.Getenv("DB_PASSWORD"),
					SSLMode:  os.Getenv("DB_SSL_MODE"),
					MinConn:  getEnvAsInt("DB_MIN_CONN", 3),
					MaxConn:  getEnvAsInt("DB_MAX_CONN", 100),
				},
			},
			Log: Log{
				LogLevel:          os.Getenv("LOG_LEVEL"),
				AddSource:         getEnvAsBool("LOG_ADD_SOURCE", false),
				ConsoleDecoration: getEnvAsBool("LOG_CONSOLE_DECORATION", true),
			},
		},
	}

	return Env, nil
}

// getEnvAsInt obtiene el string desde una variable de entorno
// y lo devuelve como int, por defecto retorna un int
func getEnvAsInt(env string, defaultVal int) int {
	valStr := os.Getenv(env)
	if valStr == "" {
		return defaultVal
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		fmt.Println("Error parsing env var", env, "as int:", err)
		return defaultVal
	}

	return val
}

// getEnvAsBoolobtiene el string desde una variable de entorno
// y lo devuelve como bool, por defecto retorna un bool
func getEnvAsBool(key string, defaultVal bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}

	parsed, err := strconv.ParseBool(val)
	if err != nil {
		return defaultVal
	}

	return parsed
}
