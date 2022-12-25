package main

import (
	"log"

	"github.com/SamMHD/simple-broker/broker"
	"github.com/SamMHD/simple-broker/destination"
	"github.com/SamMHD/simple-broker/receiver"
	"github.com/SamMHD/simple-broker/sender"
	"github.com/SamMHD/simple-broker/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	destinationServer, err := destination.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}
	go destinationServer.Start()

	brokerServer, err := broker.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}
	go brokerServer.Start()

	receiverServer, err := receiver.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}
	go receiverServer.Start()

	sender.StartSendProcedure(config)
}
