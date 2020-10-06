package environment

import (
	"errors"
	"net/http"
	"os"
	"strconv"
)

// ErrNoTargetHost occurs when no environment variable was provided
var ErrNoTargetHost = errors.New("target_host environment variable needs to be defined")

// ErrUnrecognizableAppendURI occurs when the append_uri environment variable couldn't be parsed into a bool
var ErrUnrecognizableAppendURI = errors.New("append_uri was not set or had invalid format")

// Variables hold all the configuration data
type Variables struct {
	TargetHost string `json:"target_host"`
	AppendURI  bool   `json:"append_uri"`
	LogFormat  string `json:"log_format"`
	HTTPStatus int    `json:"http_status"`
}

func Load() (v Variables, errs []error) {
	var err error

	v.TargetHost = os.Getenv("target_host")

	if v.TargetHost == "" {
		errs = append(errs, ErrNoTargetHost)
	}

	v.AppendURI, err = strconv.ParseBool(os.Getenv("append_uri"))
	if err != nil {
		errs = append(errs, ErrUnrecognizableAppendURI)
	}

	v.HTTPStatus, err = strconv.Atoi(os.Getenv("http_status"))
	if err != nil {
		v.HTTPStatus = http.StatusTemporaryRedirect
	}

	v.LogFormat = os.Getenv("log_format")

	return
}
