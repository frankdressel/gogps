package internal

import (
	"errors"
	"strconv"
	"strings"
)

func Decode(raw string, num_deg int) float64 {
	deg, _ := strconv.ParseFloat(raw[0:num_deg], 64)
	min, _ := strconv.ParseFloat(raw[num_deg:9], 64)
	return deg + min/60
}

func Parse(data string) (float64, float64, error) {
	if strings.Index(data, "$GPRMC") == 0 {
		split := strings.Split(data, ",")
		lat_raw := split[3]
		lon_raw := split[5]

		return Decode(lat_raw, 2), Decode(lon_raw, 3), nil
	}

	return -1, -1, errors.New("Only able to parseRMCs but got: " + data)
}
