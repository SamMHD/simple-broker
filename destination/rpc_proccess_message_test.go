package destination

import (
	"context"
	"testing"
	"time"

	"github.com/SamMHD/simple-broker/pb"
	"github.com/stretchr/testify/require"
)

func TestTransferMessage(t *testing.T) {

	server, err := NewServer(testConfig)

	servingError := make(chan error, 1)
	go func(servingError chan error) {
		servingError <- server.Start()
	}(servingError)

	select {
	case err := <-servingError:
		require.NoError(t, err)
	case <-time.After(1 * time.Second):
		require.NoError(t, err)
		_, err = server.ProccessMessage(context.Background(), &pb.ProccessMessageRequest{
			Message: "test",
		})
		require.NoError(t, err)
		server.Stop()
	}
}
