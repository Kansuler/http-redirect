package main

import (
	"github.com/Kansuler/http-redirect/internal/environment"
	"github.com/Kansuler/http-redirect/internal/handler"
	"github.com/Kansuler/http-redirect/internal/logging"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	config, errs := environment.Load()
	switch config.LogFormat {
	case "stackdriver":
		logging.Stackdriver()
	default:
	}

	if len(errs) > 0 {
		log.Fatal().Errs("error", errs).Msgf("error setting up required environment variables")
		return
	}

	log.Fatal().Err(http.ListenAndServe(":8080", handler.New(config))).Msgf("http listener got a fatal error")
}
