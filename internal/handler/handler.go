package handler

import (
	"fmt"
	"github.com/Kansuler/http-redirect/internal/environment"
	"net/http"
)

type handler struct {
	config environment.Variables
}

func New(config environment.Variables) http.Handler {
	return handler{
		config: config,
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.config.AppendUri {
		http.Redirect(w, r, fmt.Sprintf("%s%s", h.config.TargetHost, r.RequestURI), h.config.HttpStatus)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("%s", h.config.TargetHost), h.config.HttpStatus)
}
