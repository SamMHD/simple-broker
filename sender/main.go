// package sender will generate random messages and send them to the receiver over HTTP API
// over specified rate and duration
// message size limits and reciever address are defined in the environment variables
package sender

import (
	"sync"
	"time"

	"github.com/SamMHD/simple-broker/util"

	"github.com/rs/zerolog/log"
)

const TPS = 1000        // messages per second
const testDuration = 20 // duration in which the sender will send messages in rate of TPS

// StartSendProcedure will send messages at a rate of TPS for testDuration seconds
func StartSendProcedure(config util.Config) {
	reqCounter := 0       // counts number of sent requests
	resCounter := 0       // counts number of received responses
	var wg sync.WaitGroup // wait group to wait for all requests to finish

	// send messages for testDuration seconds
	// NOTE: the test might take more than testDuration seconds to finish
	//	     because of the time it takes to receive responses
	for start := time.Now(); time.Since(start) < time.Second*testDuration; {
		// increment the request counter
		reqCounter++
		// if the request counter is a multiple of 1000, logs the number of sent messages
		if reqCounter%1000 == 0 {
			log.Warn().Str("ser_name", "sender").Msgf("Sending %dth Request", reqCounter)
		}

		wg.Add(1)
		go send(config, &resCounter, &wg)

		// sleep for 1/TPS seconds
		time.Sleep(time.Second / TPS)
		// NOTE: this might cause the sender to send less than TPS messages per second
		//       because of the time it takes to send a message
	}

	// wait for all requests to finish
	wg.Wait()
	// log that all requests have finished
	log.Warn().Str("ser_name", "sender").Msg("Sending Requests Finished, Other Services might still be running")
	// NOTE: after this point, the sender will not send any more messages and current goroutine will exit immediately
	//       but other services might still be running (in case of building them using TargetService set to "all")
	// 	     to stop all services, use Ctrl+C or uncomment the following line
	// os.Exit(0)
}

// send will send a message to the receiver and wait for a response
// if there is an error, it will try to send the message again
// else it will retry sending the message via scheduling a new goroutine
func send(config util.Config, resCounter *int, wg *sync.WaitGroup) {
	// send an arbitrary message to the receiver
	err := sendMessage(config)
	if err != nil {
		// if there is an error, log the error and try to send the message again
		log.Debug().Str("ser_name", "sender").Msgf("%s", err)
		go send(config, resCounter, wg)
		// make sure to return to close current goroutine
		return
	}

	// if there is no error, increment the response counter and decrement the wait group
	wg.Done()
	(*resCounter)++

	// if the response counter is a multiple of 1000, logs the number of received messages
	if *resCounter%1000 == 0 {
		log.Warn().Str("ser_name", "sender").Msgf("Received %dth Response", *resCounter)
	}
}
