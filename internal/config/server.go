package config

import (
	"fmt"
	"net/http"
	"time"

)

func StartServer(mux *http.ServeMux, port string) error {
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
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
