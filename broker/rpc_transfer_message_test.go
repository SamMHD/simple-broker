package broker

import (
	"context"
	"testing"

	"github.com/SamMHD/simple-broker/pb"
	"github.com/stretchr/testify/require"
)

func TestTransferMessage(t *testing.T) {

	prepareDestinationServiceForTest(testConfig, t)
	server, err := NewServer(testConfig)
	require.NoError(t, err)
	_, err = server.TransferMessage(context.Background(), &pb.TransferMessageRequest{
		Message: "test",
	})
	require.NoError(t, err)
}
