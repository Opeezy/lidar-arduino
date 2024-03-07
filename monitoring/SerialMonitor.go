package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

func main() {
	serialConfig := &serial.Config{
		Name:        "COM5",
		Baud:        9600,
		Size:        8,
		StopBits:    1,
		ReadTimeout: time.Second * 5}

	serialPort, err := serial.OpenPort(serialConfig)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 8)
	i, err := serialPort.Read(buf)
	_ = i
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range buf {
		fmt.Printf("%x", l)
	}
	fmt.Print("\n")

}
