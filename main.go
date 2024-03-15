package main

import (
	"log"
	"time"
)

func main() {
	// Initialize our Serial component
	timeOut := time.Second * 5
	newSerial := Serial{
		configMade: false,
		portOpen:   false,
	}
	newSerial.makeConfig("COM5", 230400, 8, 1, timeOut)

	// Check for valid config and attempt to open serial port connection
	err := newSerial.openPortConnection()
	if err != nil {
		log.Fatalln(err)
	} else {
		// If everything goes well, read incoming serial data...
		newSerial.listen(time.Second * 10)
	}
}
