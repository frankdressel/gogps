package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/tarm/serial"
)

func decode(raw string, num_deg int) float64 {
	deg, _ := strconv.ParseFloat(raw[0:num_deg], 64)
	min, _ := strconv.ParseFloat(raw[num_deg:9], 64)
	return deg + min/60
}

func parse(data string) {
	if strings.Index(data, "$GPRMC") == 0 {
		split := strings.Split(data, ",")
		lat_raw := split[3]
		lon_raw := split[5]
		fmt.Println(decode(lat_raw, 2), decode(lon_raw, 3))
	}
}

func main() {
	c := &serial.Config{Name: "/dev/ttyAMA1", Baud: 9600}
	s, err := serial.OpenPort(c)
	defer s.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		data := scanner.Text()
		parse(data)
	}

}
