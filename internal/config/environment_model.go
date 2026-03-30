package config

type (
	env struct {
		Container Container
	}

	Container struct {
		App App
		DB  DB
		JWT JWT
		Log Log
	}

	App struct {
		Name           string
		Protocol       string
		Port           string
		AllowedOrigins string
		AllowedMethods string
	}

	DB struct {
		PostgresEnv PostgresEnv
	}

	PostgresEnv struct {
		Host     string
		Port     string
		Name     string
		User     string
		Password string
		SSLMode  string
		MinConn  int
		MaxConn  int
	}

	JWT struct {
		ExpirationTime int64
		SecretKey      string
	}

	Log struct {
		LogLevel          string
		AddSource         bool
		ConsoleDecoration bool
	}
)
