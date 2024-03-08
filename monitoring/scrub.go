package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

func split(data string) [][]string {
	var packet [][]string
	var pack []string
	count := 0

	for _, v := range data {
		if count > 1 {
			pack = append(pack, string(v))
			count++
		} else {
			packet = append(packet, pack)
			count = 0
		}
	}
	return packet
}

func main() {
	serialConfig := &serial.Config{
		Name:        "COM5",
		Baud:        230400,
		Size:        8,
		StopBits:    1,
		ReadTimeout: time.Second * 5}

	serialPort, err := serial.OpenPort(serialConfig)
	if err != nil {
		log.Fatal(err)
	}

	for {
		buf := make([]byte, 100)

		i, _ := serialPort.Read(buf)
		_ = i

		for _, v := range buf {
			fmt.Printf("%x ", v)
		}
	}
}
