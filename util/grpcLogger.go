package util

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewGrpcLoggerForService returns a new gRPC server interceptor that logs the execution of each gRPC method.
// It uses service name to create a logger with the service name as a field.
// return result is a function which can be used as a gRPC server interceptor.
func NewGrpcLoggerForService(serviceName string) func(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {

	// returning new anonymous function which can be used as a gRPC server interceptor
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		// recording the start time of the request in order to calculate the response time
		startTime := time.Now()

		// calling the next handler to handle the request
		result, err := handler(ctx, req)

		// calculating the response time using startTime checkpoint
		duration := time.Since(startTime)

		// using codes.Unknown as the default status code inorder to avoid nil pointer dereference
		statusCode := codes.Unknown
		// convert error to status code if it is convertable
		if st, ok := status.FromError(err); ok {
			statusCode = st.Code()
		}

		// create an info level logger to log accessed procedure
		logger := log.Info()
		// check if there is an error and replace logger with an error level logger if there is an error
		// and wrap the error with the logger
		if err != nil {
			logger = log.Error().Err(err)
		}

		// logging the accessed procedure
		logger.Str("protocol", "grpc").
			Str("method", info.FullMethod).            // FullMethod is the full RPC method string, i.e., /package.service/method.
			Str("ser_name", serviceName).              // service name (describing the service being called)
			Int("status", int(statusCode)).            // status code of the response
			Str("status_reason", statusCode.String()). // human readable status code
			Dur("resp_time", duration).                // response time in milliseconds
			Msg("gRPC Request")

		// return the gRPC return value and the error as the result of the interceptor
		return result, err
	}
}
