package environment

import (
	"errors"
	"net/http"
	"os"
	"strconv"
)

var ErrNoTargetHost = errors.New("target_host environment variable needs to be defined")
var ErrUnrecognizableAppendUri = errors.New("append_uri was not set or had invalid format")

type Variables struct {
	TargetHost string `json:"target_host"`
	AppendUri  bool   `json:"append_uri"`
	LogFormat  string `json:"log_format"`
	HttpStatus int    `json:"http_status"`
}

func Load() (v Variables, errs []error) {
	var err error

	v.TargetHost = os.Getenv("target_host")

	if v.TargetHost == "" {
		errs = append(errs, ErrNoTargetHost)
	}

	v.AppendUri, err = strconv.ParseBool(os.Getenv("append_uri"))
	if err != nil {
		errs = append(errs, ErrUnrecognizableAppendUri)
	}

	v.HttpStatus, err = strconv.Atoi(os.Getenv("http_status"))
	if err != nil {
		v.HttpStatus = http.StatusTemporaryRedirect
	}

	v.LogFormat = os.Getenv("log_format")

	return
}
