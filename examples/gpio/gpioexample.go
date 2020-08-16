package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/balazsgrill/phirmata"
	"github.com/tarm/serial"
	"periph.io/x/periph/conn/gpio"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "COM1", "Serial port")
	flag.Parse()

	fmt.Printf("Opening %s\n", port)

	conn, err := serial.OpenPort(&serial.Config{Name: port, Baud: 57600})
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	dev, err := phirmata.New(conn)
	if err != nil {
		panic(err)
	}

	t := time.NewTicker(time.Second)
	v := gpio.High
	for range t.C {
		err = dev.Pins()[13].Out(v)
		if err != nil {
			panic(err)
		}
		if v == gpio.High {
			v = gpio.Low
		} else {
			v = gpio.High
		}
	}
}
