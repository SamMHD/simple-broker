// package sender will generate random messages and send them to the receiver over HTTP API
// over specified rate and duration
// message size limits and reciever address are defined in the environment variables
package sender

import (
	"time"

	"github.com/SamMHD/simple-broker/util"

	"github.com/rs/zerolog/log"
)

const TPS = 10000
const testDuration = 10

// StartSendProcedure will send messages at a rate of TPS for testDuration seconds

func StartSendProcedure(config util.Config) {
	// send messages at a rate of TPS for testDuration seconds
	i := 0
	for start := time.Now(); time.Since(start) < time.Second*testDuration; {
		// send a message
		log.Info().Msgf("sending message %d", i)
		i++
		err := sendMessage(config)
		if err != nil {
			log.Fatal().Msgf("%s", err)
		}

		// sleep for 1/TPS seconds
		time.Sleep(time.Second / TPS)
	}
}
