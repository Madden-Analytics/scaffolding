package main

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn())
	defer pool.Close()
	if err != nil {
		log.Fatal().Err(err).Msg("could not create the database pool")
	}

	startupCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(startupCtx); err != nil {
		log.Fatal().Err(err).Msg("could not reach the database")
	}
	log.Info().Msg("connected to the database")

	r := chi.NewRouter()

	// The frontend dev server runs on another origin, so it needs CORS to call us.
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{env("CORS_ORIGIN", "http://localhost:5173")},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		MaxAge:         300,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server running OK"))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		pingCtx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		code := http.StatusOK
		if err := pool.Ping(pingCtx); err != nil {
			log.Error().Err(err).Msg("health check failed")
			code = http.StatusServiceUnavailable
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"status": http.StatusText(code)})
	})

	log.Info().Msg("server starting on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal().Err(err).Msg("server stopped")
	}
}

// env returns the value of the environment variable, or fallback when it is unset.
func env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// dsn builds the Postgres connection string from the environment. It uses net/url
// so that passwords containing characters like @, / or # are escaped correctly.
func dsn() string {
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(env("DB_USER", "postgres"), env("DB_PASSWORD", "postgres")),
		Host:     net.JoinHostPort(env("DB_HOST", "localhost"), env("DB_PORT", "5432")),
		Path:     "/" + env("DB_NAME", "main_db"),
		RawQuery: "sslmode=disable",
	}
	return u.String()
}
