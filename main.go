package main

import (
	"log"
	"time"
)

func main() {
	// Initialize our Serial component
	timeOut := time.Second * 5
	newSerial := Serial{
		ConfigMade: false,
		PortOpen:   false,
	}
	newSerial.MakeConfig("COM5", 230400, 8, 1, timeOut)

	// Check for valid config and attempt to open serial port connection
	err := newSerial.OpenPortConnection()
	if err != nil {
		log.Fatalln(err)
	} else {
		// If everything goes well, read incoming serial data...
		newSerial.Listen(time.Second * 10)
	}
}
