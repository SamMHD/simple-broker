// package sender will generate random messages and send them to the receiver over HTTP API
// over specified rate and duration
// message size limits and reciever address are defined in the environment variables
package sender

import (
	"log"
	"time"

	"github.com/SamMHD/simple-broker/util"
)

const TPS = 10000
const testDuration = 10

var config util.Config

func init() {
	var err error
	config, err = util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
}

// StartSendProcedure will send messages at a rate of TPS for testDuration seconds
func StartSendProcedure() {
	// send messages at a rate of TPS for testDuration seconds
	for start := time.Now(); time.Since(start) < time.Second*testDuration; {
		// send a message
		err := sendMessage(config)
		if err != nil {
			log.Fatal(err)
		}

		// sleep for 1/TPS seconds
		time.Sleep(time.Second / TPS)
	}
}
