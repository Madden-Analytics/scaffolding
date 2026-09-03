package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// env returns the value of the environment variable, or fallback when it is unset.
func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env("DB_USER", "postgres"),
		env("DB_PASSWORD", "postgres"),
		env("DB_HOST", "localhost"),
		env("DB_PORT", "5432"),
		env("DB_NAME", "main_db"),
	)

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("could not create the database pool")
	}
	defer pool.Close()

	startupCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(startupCtx); err != nil {
		log.Fatal().Err(err).Msg("could not reach the database")
	}
	log.Info().Msg("connected to the database")

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server running OK"))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		pingCtx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		status, code := "ok", http.StatusOK
		if err := pool.Ping(pingCtx); err != nil {
			log.Error().Err(err).Msg("health check failed")
			status, code = "unavailable", http.StatusServiceUnavailable
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"status": status})
	})

	log.Info().Msg("server starting on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal().Err(err).Msg("server stopped")
	}
}
