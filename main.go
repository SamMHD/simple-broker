package main

import (
	"log"

	"github.com/SamMHD/simple-broker/receiver"
	"github.com/SamMHD/simple-broker/sender"
	"github.com/SamMHD/simple-broker/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	receiverServer, err := receiver.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}
	go receiverServer.Start()

	sender.StartSendProcedure(config)
}
