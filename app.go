package main

import (
	"log"
	"time"

	"github.com/Opeezy/utils"
)

func main() {
	// Initialize our Serial component
	Serial1 := utils.NewSerial("COM5", 230400, 8, 1, time.Second*5)

	// Check for valid config and attempt to open serial port connection
	err := Serial1.OpenPortConnection()
	if err != nil {
		log.Fatalln(err)
	} else {
		// If everything goes well, read incoming serial data...
		Serial1.Listen(time.Second * 10)
	}
}
