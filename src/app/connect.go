package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

type Serial struct {
	name       string
	config     *serial.Config
	port       *serial.Port
	writeReady byte
	readReady  byte
	configMade bool
	portOpen   bool
	connStatus bool
}

type SerialWorker interface {
	openPortConnection()
	makeConfig()
	handleCode()
}

func (s *Serial) openPortConnection() error {
	log.Printf("Opening port to %v\n", s.name)
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

func (s *Serial) makeConfig() {
	newConfig := &serial.Config{
		Name:        "COM5",
		Baud:        230400,
		Size:        8,
		StopBits:    1,
		ReadTimeout: time.Second * 5}
	s.config = newConfig
	s.name = newConfig.Name
	s.configMade = true
	log.Println("Config generated")
}

func (s *Serial) read() {

	i, _ := s.port.Write([]byte{s.readReady})
	_ = i

}

func (s *Serial) handleCode(code byte) {
	switch code {
	case s.writeReady:
	}
}

func main() {

	newSerial := Serial{
		writeReady: 17,
		readReady:  18,
		configMade: false,
		portOpen:   false,
		connStatus: false,
	}

	newSerial.makeConfig()

	for !newSerial.portOpen {
		err := newSerial.openPortConnection()
		if err != nil {
			log.Println(err)
			log.Printf("%v not open", newSerial.name)
		} else {
			break
		}
		time.Sleep(time.Second * 2)
	}

	for {
		fmt.Println("Enter code: ")
		code, err := fmt.Scanln()

		if err != nil {
			log.Println(err)
		}

		newSerial.handleCode(byte(code))
	}
}
