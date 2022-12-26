// package sender will generate random messages and send them to the receiver over HTTP API
// over specified rate and duration
// message size limits and reciever address are defined in the environment variables
package sender

import (
	"fmt"
	"sync"
	"time"

	"github.com/SamMHD/simple-broker/util"

	"github.com/rs/zerolog/log"
)

const TPS = 10000
const testDuration = 2

// StartSendProcedure will send messages at a rate of TPS for testDuration seconds

func StartSendProcedure(config util.Config) {
	// send messages at a rate of TPS for testDuration seconds
	i := 0

	// new wait group
	var wg sync.WaitGroup
	f := 0

	for start := time.Now(); time.Since(start) < time.Second*testDuration; {
		// send a message
		log.Info().Msgf("sending message %d", i)
		i++
		if i%1000 == 0 {
			fmt.Println("Performing", i, "th message")
		}

		wg.Add(1)
		go func() {
			err := sendMessage(config)
			if err != nil {
				fmt.Println("error")
				log.Error().Str("ser_name", "sender").Msgf("%s", err)
			}
			wg.Done()
			f++
			if f%1000 == 0 {
				fmt.Println("Received", f, "th message")
			}
		}()

		// sleep for 1/TPS seconds
		time.Sleep(time.Second / TPS)
	}
	fmt.Println("here")
	wg.Wait()
}
