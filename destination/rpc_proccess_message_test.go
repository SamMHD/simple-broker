package destination

import (
	"context"
	"testing"

	"github.com/SamMHD/simple-broker/pb"
	"github.com/stretchr/testify/require"
)

func TestTransferMessage(t *testing.T) {

	server, err := NewServer(testConfig)
	require.NoError(t, err)
	_, err = server.ProccessMessage(context.Background(), &pb.ProccessMessageRequest{
		Message: "test",
	})
	require.NoError(t, err)
}
