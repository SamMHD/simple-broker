package main

import (
	"github.com/SamMHD/simple-broker/receiver"
	"github.com/SamMHD/simple-broker/sender"
)

func main() {
	receiver.StartListening()
	sender.StartSendProcedure()
}
