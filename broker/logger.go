package broker

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

// newLogger creates a new zerolog logger and returns it.
// destination can be "stdout", "stderr" or a file path.
// (if destination is not stdout/stderr it assumes that it's a filepath)
// while using file path as destination, it creates the file if it doesn't exist
// and in case of any error it will use stdout as destination.
// it also returns an error if any.
func newLogger(destination string) (zerolog.Logger, error) {
	if destination == "stdout" {
		// create a new logger with stdout as destination
		return zerolog.New(os.Stdout).Level(zerolog.InfoLevel), nil
	} else if destination == "stderr" {
		// create a new logger with stderr as destination
		return zerolog.New(os.Stderr).Level(zerolog.InfoLevel), nil
	}

	// open the file or create it if it doesn't exist
	file, err := os.OpenFile(
		destination,                         // file path (can be relative or absolute)
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, // append to the file if it exists, create it if it doesn't exist and open it in write only mode
		0664,                                // rw-rw-r-- premission
	)

	if err != nil {
		// if there is an error while opening the file, use stdout as destination
		// and return the error
		return zerolog.New(os.Stdout).Level(zerolog.InfoLevel),
			fmt.Errorf("logging to file failed! (using stdout instead) Error: %s", err)
	}

	// if there is no error, create a new logger with the file as destination
	return zerolog.New(file).With().Timestamp().Logger(), nil
}
