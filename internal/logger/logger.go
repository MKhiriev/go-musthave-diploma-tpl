package logger

import (
	"os"
	"runtime"

	"github.com/rs/zerolog"
)

func NewLogger(role string) *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return runtime.FuncForPC(pc).Name() // return function name
	}

	zerolog.CallerFieldName = "func"
	logger := zerolog.New(os.Stdout).With().
		Str("role", role).
		Timestamp().
		Caller().
		Logger()

	return &logger
}
