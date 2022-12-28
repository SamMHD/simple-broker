package broker

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	testCases := []struct {
		name    string
		prepare func() string
		judge   func(zerolog.Logger, error) error
		cleanup func()
	}{
		{
			name: "destiantion_stdout",
			prepare: func() string {
				return "stdout"
			},
			judge: func(logger zerolog.Logger, err error) error {
				if err != nil {
					return err
				}
				if logger.GetLevel() != zerolog.InfoLevel {
					return fmt.Errorf("expected level to be %s, got %s", zerolog.InfoLevel, logger.GetLevel())
				}
				return nil
			},
			cleanup: func() {},
		},
		{
			name: "destiantion_stderr",
			prepare: func() string {
				return "stderr"
			},
			judge: func(logger zerolog.Logger, err error) error {
				if err != nil {
					return err
				}
				if logger.GetLevel() != zerolog.InfoLevel {
					return fmt.Errorf("expected level to be %s, got %s", zerolog.InfoLevel, logger.GetLevel())
				}
				return nil
			},
			cleanup: func() {},
		},
		{
			name: "destiantion_file",
			prepare: func() string {
				return "test.log"
			},
			judge: func(logger zerolog.Logger, err error) error {
				if err != nil {
					return err
				}
				if logger.GetLevel() != zerolog.InfoLevel {
					return fmt.Errorf("expected level to be %s, got %s", zerolog.InfoLevel, logger.GetLevel())
				}
				return nil
			},
			cleanup: func() {
				// delete the file
				os.Remove("test.log")
			},
		},
		{
			name: "destiantion_file_error",
			prepare: func() string {
				return "/test.log"
			},
			judge: func(logger zerolog.Logger, err error) error {
				if err == nil {
					return fmt.Errorf("expected error, got nil")
				}
				return nil
			},
			cleanup: func() {},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dest := tc.prepare()
			returned_logger, returned_err := newLogger(dest)
			err := tc.judge(returned_logger, returned_err)
			require.NoError(t, err)
			tc.cleanup()
		})
	}
}
