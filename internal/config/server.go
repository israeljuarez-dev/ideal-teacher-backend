package config

import (
	"fmt"
	"net/http"
	"time"

)

func StartServer(handler http.Handler, port string) error {
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Printf("🚀 Servidor [Mux] corriendo en http://localhost:%s\n", port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("fallo al iniciar el servidor: %w", err)
	}

	return nil
}
