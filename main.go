// package main contains all of the services that are used to run the application
// it also contains the main function that starts the application
// you can specify which services to run by setting the Mode variable (all, destination, broker, receiver, sender)
// it can also be set using build time LDFlags (see Makefile)
package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/SamMHD/simple-broker/broker"
	"github.com/SamMHD/simple-broker/destination"
	"github.com/SamMHD/simple-broker/receiver"
	"github.com/SamMHD/simple-broker/sender"
	"github.com/SamMHD/simple-broker/util"

	"github.com/rs/zerolog/log"
)

var TargetService string           // variable used to specify which services to run (configured using LDFlags)
var GlobalWaitGroup sync.WaitGroup // global wait group used to wait for all services to finish

func main() {
	// defining command line flags
	// noStartUpMessage: disable startup message
	noStartUpMessage := flag.Bool("no-startup-message", false, "disable startup message")
	// logType: log type (console, json)
	logType := flag.String("log-type", "console", "stderr log type (console, json)")
	// logLevel: log level (debug, info, warn, error)
	logLevel := flag.String("log-level", "info", "stderr log level (debug, info, warn, error)")

	// parse flags
	flag.Parse()

	// startup message
	if !*noStartUpMessage {
		fmt.Printf("Starting Simple Broker(%s)...\n", TargetService)
	}

	// set log settings
	util.SetGlobalLogSettings(*logType, *logLevel)

	// load config file
	config, err := util.LoadConfig(".")
	if err != nil {
		// in case of error, log the error and exit
		log.Fatal().Err(err).Msg("failed to load config file")
	}

	// start services...

	// start destination service if TargetService is "all" or "destination"
	if TargetService == "all" || TargetService == "destination" {
		// create destination service server
		destinationServer, err := destination.NewServer(config)
		if err != nil {
			// in case of error, log the error and exit
			log.Fatal().Err(err).Msg("failed to create destination service server")
		}

		// scheduling a new goroutine to start the server
		GlobalWaitGroup.Add(1)
		go func() {
			// start the server or catch the error if it fails
			err := destinationServer.Start()
			if err != nil {
				// in case of error, log the error and exit
				log.Fatal().Err(err).Msg("failed to start destination service server")
			}
			// in case of peacefull exit, decrement the wait group
			GlobalWaitGroup.Done()
		}()
	}

	// start broker service if TargetService is "all" or "broker"
	if TargetService == "all" || TargetService == "broker" {
		// create broker service server
		brokerServer, err := broker.NewServer(config)
		if err != nil {
			// in case of error, log the error and exit
			log.Fatal().Err(err).Msg("failed to create broker service server")
		}

		// scheduling a new goroutine to start the server
		GlobalWaitGroup.Add(1)
		go func() {
			// start the server or catch the error if it fails
			err := brokerServer.Start()
			if err != nil {
				// in case of error, log the error and exit
				log.Fatal().Err(err).Msg("failed to start broker server")
			}
			// in case of peacefull exit, decrement the wait group
			GlobalWaitGroup.Done()
		}()
	}

	// start receiver service if TargetService is "all" or "receiver"
	if TargetService == "all" || TargetService == "receiver" {
		// create receiver service server
		receiverServer, err := receiver.NewServer(config)
		if err != nil {
			// in case of error, log the error and exit
			log.Fatal().Err(err).Msg("failed to create receiver service server")
		}

		// scheduling a new goroutine to start the server
		GlobalWaitGroup.Add(1)
		go func() {
			// start the server or catch the error if it fails
			err := receiverServer.Start()
			if err != nil {
				// in case of error, log the error and exit
				log.Fatal().Err(err).Msg("failed to start receiver server")
			}
			// in case of peacefull exit, decrement the wait group
			GlobalWaitGroup.Done()
		}()
	}

	// start sender service if TargetService is "all" or "sender"
	if TargetService == "all" || TargetService == "sender" {
		// scheduling a new goroutine to start the sender
		GlobalWaitGroup.Add(1)
		go func() {
			// start sending procedure
			sender.StartSendProcedure(config)
			// decrement the wait group
			GlobalWaitGroup.Done()
		}()
	}

	// wait for all services to finish
	// they might not finish, but in case of peacefull exit, they will.
	GlobalWaitGroup.Wait()
}
