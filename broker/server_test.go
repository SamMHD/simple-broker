package broker

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/SamMHD/simple-broker/destination"
	"github.com/SamMHD/simple-broker/util"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/require"
)

var testDestinationServiceRunning sync.Mutex

func prepareDestinationServiceForTest(config util.Config, t *testing.T) context.CancelFunc {
	testDestinationServiceRunning.Lock()
	ctx, cancel := context.WithCancel(context.Background())

	server, err := destination.NewServer(config)
	require.NoError(t, err)

	safeCancel := context.CancelFunc(func() {
		if ctx.Err() == nil {
			cancel()
		}
	})

	go func() {
		err := server.Start()
		safeCancel()
		require.NoError(t, err)
	}()

	go func() {
		<-ctx.Done()
		server.Stop()
		testDestinationServiceRunning.Unlock()
	}()

	return safeCancel
}

func setConfigPorts(config util.Config) util.Config {
	brokerPort, _ := freeport.GetFreePort()
	config.BrokerAddress = strings.ReplaceAll(config.BrokerAddress, "<PORT>", fmt.Sprint(brokerPort))

	desPort, _ := freeport.GetFreePort()
	config.DestinationAddress = strings.ReplaceAll(config.DestinationAddress, "<PORT>", fmt.Sprint(desPort))

	return config
}

var rawTestConfig util.Config = util.Config{
	DestinationAddress:   "localhost:<PORT>",
	BrokerAddress:        "localhost:<PORT>",
	BrokerLogDestination: "./test.log",
}

func TestNewBrokerServer(t *testing.T) {
	testConfig := setConfigPorts(rawTestConfig)
	stopDesServer := prepareDestinationServiceForTest(testConfig, t)
	defer stopDesServer()

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
			name: "invalid_broker_log_destination",
			prepare: func() util.Config {
				config := testConfig
				config.BrokerLogDestination = "/tmp"
				return config
			},
			judge: func(server *Server, err error) error {
				if err == nil {
					return fmt.Errorf("expected error to be not nil")
				}
				if server != nil {
					return fmt.Errorf("expected server to be nil")
				}
				return nil
			},
			cleanup: func() {},
		},
		{
			name: "invalid_broker_address_port_busy",
			prepare: func() util.Config {
				config := testConfig
				config.BrokerAddress = config.DestinationAddress
				return config
			},
			judge: func(server *Server, err error) error {
				servingError := make(chan error, 1)

				go func(servingError chan error) {
					servingError <- server.Start()
				}(servingError)

				select {
				case err := <-servingError:
					if err == nil {
						return fmt.Errorf("expected error to be not nil")
					}
					return nil
				case <-time.After(5 * time.Second):
					server.Stop()
					return fmt.Errorf("expected server to face error in 5 seconds")
				}
			},
			cleanup: func() {},
		},
		{
			name: "invalid_destination_address",
			prepare: func() util.Config {
				config := setConfigPorts(rawTestConfig)
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

	os.Remove("./test.log")
}
