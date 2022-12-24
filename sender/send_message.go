package sender

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/SamMHD/simple-broker/util"
	"github.com/bxcodec/faker/v3"
)

// sendMessage will generate a random message and send it to the receiver
// message size limits are defined in the environment variable MESSAGE_MIN_SIZE and MESSAGE_MAX_SIZE
// reciever address is defined in the environment variable RECEIVER_ADDRESS and RECEIVER_PORT
func sendMessage(config util.Config) error {
	// generate a random message
	message, err := generateRandomMessage(config.MessageMinSize, config.MessageMaxSize)
	if err != nil {
		return err
	}

	// send the message to the receiver over HTTP API
	err = sendHTTPMessage(config.ReceiverAddress, config.ReceiverPort, message)
	if err != nil {
		return err
	}

	return nil
}

// generateRandomMessage generate a random english message of size between minSize and maxSize
// using faker library to generate random integers and words
func generateRandomMessage(minSize int, maxSize int) (message string, err error) {
	// generate a random message length between minSize and maxSize
	messageLengthSlice, err := faker.RandomInt(minSize, maxSize)
	if err != nil {
		return
	}
	// faker.RandomInt returns a slice of int, so we need to get the first element
	messageLength := messageLengthSlice[0]

	// generate a random message of size messageLength
	for len(message) < messageLength {
		message += faker.Sentence()
	}
	return message[0:messageLength], nil
}

// messageRequest is the request body of the HTTP API
type messageRequest struct {
	message string `json:message`
}

// sendHTTPMessage encodes message into JSON and send it to the receiver over HTTP API
func sendHTTPMessage(receiverAddress string, receiverPort string, message string) error {
	// create the request body
	requestBody, err := json.Marshal(messageRequest{message: message})
	if err != nil {
		return err
	}

	// send the request
	bodyReader := bytes.NewReader(requestBody)
	_, err = http.Post("http://"+receiverAddress+":"+receiverPort+"/message", "application/json", bodyReader)

	return nil
}
