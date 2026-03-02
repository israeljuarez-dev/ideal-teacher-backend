package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Container struct {
	App  *App
	DB   *DB
	HTTP *HTTP
}

type App struct {
	Name string
	Env  string
}

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	MinConn  int
	MaxConn  int
}

type HTTP struct {
	Env            string
	URL            string
	Port           string
	AllowedOrigins string
}

func New() (*Container, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error cargando .env: %v", err)
	}

	app := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	db := &DB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
		MinConn:  getEnvAsInt("DB_MIN_CONN", 3),
		MaxConn:  getEnvAsInt("DB_MAX_CONN", 100),
	}

	http := &HTTP{
		Env:            os.Getenv("APP_ENV"),
		URL:            os.Getenv("HTTP_URL"),
		Port:           os.Getenv("SERVER_PORT"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
	}

	return &Container{
		App:  app,
		DB:   db,
		HTTP: http,
	}, nil
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
