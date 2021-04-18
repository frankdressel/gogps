package main

import (
	"fmt"

	"github.com/frankdressel/gogps/internal"
)

func main() {
	latlonchannel := make(chan internal.LatLon)

	internal.Read(latlonchannel, "/dev/ttyAMA1", 9600)

	for l := range latlonchannel {
		fmt.Println(l)
	}
}
