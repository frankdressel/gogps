package internal

import (
	"bufio"
	"log"
	"strconv"

	"github.com/tarm/serial"
)

type LatLon struct {
	Lat float64
	Lon float64
}

func (l LatLon) String() string {
    return "{\"Lat\": " + strconv.FormatFloat(l.Lat, 'f', 4, 64) + ", \"Lon\": " + strconv.FormatFloat(l.Lon, 'f', 4, 64) + "}"
}

func read(channel chan LatLon, device string, baud int) {
	c := &serial.Config{Name: device, Baud: baud}
	s, err := serial.OpenPort(c)
	defer s.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		data := scanner.Text()
		lat, lon, err := Parse(data)
		if err == nil {
			channel <- LatLon{lat, lon}
		}
	}
}

func Read(channel chan LatLon, device string, baud int) {
	go read(channel, device, baud)
}
