package cors

import (
	"gateway/config"
	"net/http"

	"github.com/rs/cors"
)

func NewCors(cfg *config.Config) *cors.Cors {
	var debug bool
	if cfg.Server.Mode == "production" {
		debug = false
	} else {
		debug = true
	}

	options := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete},
		// AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization", "X-Forwarded-For", "X-Real-IP", "X-Requested-With"},
		AllowedHeaders: []string{"*"},

		// Enable Debugging for testing, consider disabling in production
		Debug: debug,
	}

	return cors.New(options)
}
