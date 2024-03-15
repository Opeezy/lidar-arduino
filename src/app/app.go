package main

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

type Serial struct {
	name       string         // Serial port name (eg. "COM5")
	config     *serial.Config // Serial config pointer
	port       *serial.Port   // Serial port pointer
	configMade bool           // Valid config status
	portOpen   bool           // Port connection status
}

type SerialWorker interface {
	openPortConnection()
	makeConfig()
	listen()
}

func (s *Serial) openPortConnection() error {
	port, err := serial.OpenPort(s.config)
	s.port = port
	if err != nil {
		return err
	} else {
		log.Printf("Successfully opened port %v", s.name)
		s.portOpen = true
		return nil
	}
}

func (s *Serial) makeConfig(name string, baud int, size byte, stopbits serial.StopBits, timeout time.Duration) {
	newConfig := &serial.Config{
		Name:        name,            // Port name (eg. "COM5")
		Baud:        baud,            // Baudrate
		Size:        size,            // Data size (usually 8 bytes)
		StopBits:    stopbits,        // Stop bits (usually 1)
		ReadTimeout: time.Second * 5} // Timeout duration
	s.config = newConfig
	s.name = newConfig.Name
	s.configMade = true
	log.Println("Config generated")
}

func (s *Serial) listen(duration time.Duration) {
	if !s.configMade {
		log.Fatalln("No config made")
	} else if !s.portOpen {
		log.Fatalln("No port connection")
	} else {
		// Get the current time
		timeNow := time.Now()

		for {
			// Check for duration and read serial data until it is reached
			if time.Since(timeNow) >= duration {
				log.Fatalln("5 seconds passed")
			}
		}
	}
}

func main() {
	// Initialize our Serial component
	timeOut := time.Second * 5
	newSerial := Serial{
		configMade: false,
		portOpen:   false,
	}
	newSerial.makeConfig("COM5", 230400, 8, 1, timeOut)

	// Check for valid config and attempt to open serial port connection
	if newSerial.configMade {
		err := newSerial.openPortConnection()
		if err != nil {
			log.Fatalln(err)
		} else {
			// If everything goes well, read incoming serial data...
			newSerial.listen(time.Second * 5)
		}
	}
}
