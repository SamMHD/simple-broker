package main

import (
	"os"

	"github.com/SamMHD/simple-broker/broker"
	"github.com/SamMHD/simple-broker/destination"
	"github.com/SamMHD/simple-broker/receiver"
	"github.com/SamMHD/simple-broker/sender"
	"github.com/SamMHD/simple-broker/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msgf("failed to load config file: %s", err)
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	destinationServer, err := destination.NewServer(config)
	if err != nil {
		log.Fatal().Msgf("%s", err)
	}
	go destinationServer.Start()

	brokerServer, err := broker.NewServer(config)
	if err != nil {
		log.Fatal().Msgf("%s", err)
	}
	go brokerServer.Start()

	receiverServer, err := receiver.NewServer(config)
	if err != nil {
		log.Fatal().Msgf("%s", err)
	}
	go receiverServer.Start()

	sender.StartSendProcedure(config)
}
