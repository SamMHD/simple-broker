package destination

import (
	"fmt"
	"testing"
	"time"

	"github.com/SamMHD/simple-broker/util"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/require"
)

var testConfig util.Config = util.Config{
	DestinationAddress: fmt.Sprintf("localhost:%d", freeport.GetPort()),
}

func TestNewDestinationServer(t *testing.T) {

	testCases := []struct {
		name    string
		prepare func() util.Config
		judge   func(*Server, error) error
		cleanup func()
	}{
		{
			name: "normal",
			prepare: func() util.Config {
				return testConfig
			},
			judge: func(server *Server, err error) error {
				if err != nil {
					return err
				}
				if server == nil {
					return fmt.Errorf("expected server to be not nil")
				}

				servingError := make(chan error, 1)
				go func(servingError chan error) {
					servingError <- server.Start()
				}(servingError)

				select {
				case err := <-servingError:
					return err
				case <-time.After(2 * time.Second):
					server.Stop()
					return nil
				}
			},
			cleanup: func() {},
		},
		{
			name: "invalid_destination_address",
			prepare: func() util.Config {
				config := testConfig
				config.DestinationAddress = "test"
				return config
			},
			judge: func(server *Server, err error) error {
				servingError := make(chan error, 1)
				go func(serverError chan error) {
					servingError <- server.Start()
				}(servingError)

				select {
				case err := <-servingError:
					if err == nil {
						return fmt.Errorf("expected server to face")
					}
					return nil
				case <-time.After(5 * time.Second):
					server.Stop()
					return fmt.Errorf("expected server to face error in 5 seconds")
				}
			},
			cleanup: func() {},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			config := testCase.prepare()
			server, err := NewServer(config)
			err = testCase.judge(server, err)
			require.NoError(t, err)
			testCase.cleanup()
		})
	}
}
