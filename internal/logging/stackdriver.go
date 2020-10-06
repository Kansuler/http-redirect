package logging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"time"
)

func Stackdriver() {
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		if l == zerolog.FatalLevel {
			return "CRITICAL"
		}

		return strings.ToUpper(l.String())
	}

	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = zerolog.New(os.Stdout)
}
