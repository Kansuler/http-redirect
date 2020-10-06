package handler

import (
	"fmt"
	"github.com/Kansuler/http-redirect/internal/environment"
	"net/http"
)

type handler struct {
	config environment.Variables
}

// New creates a custom http handler that uses configuration to decide what to do with the request
func New(config environment.Variables) http.Handler {
	return handler{
		config: config,
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.config.AppendURI {
		http.Redirect(w, r, fmt.Sprintf("%s%s", h.config.TargetHost, r.RequestURI), h.config.HTTPStatus)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("%s", h.config.TargetHost), h.config.HTTPStatus)
}
