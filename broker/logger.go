package broker

import (
	"os"

	"github.com/rs/zerolog"
)

func NewLogger() zerolog.Logger {
	Logger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)
	return Logger
}
