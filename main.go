package main

import (
	"fmt"
	"net/http"

	"github.com/frankdressel/gogps/internal"
	"github.com/gorilla/websocket"

	"github.com/rs/zerolog/log"
)

type Client struct {
	conn *websocket.Conn
}

var upgrader = websocket.Upgrader{}
var clients = make(map[Client]bool)

func broadcast(message string) {
	for c, _ := range clients {
		w, err := c.conn.NextWriter(websocket.TextMessage)
		if err == nil {
			w.Write([]byte(message))
		} else {
			delete(clients, c)
			log.Error().Msgf("Error while sending data to cllient: %s", err.Error())
		}
	}
}

func gps(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	cl := Client{c}
	clients[cl] = true
}

func main() {

	go func() {
		var latlonchannel = make(chan internal.LatLon)
		internal.Read(latlonchannel, "/dev/ttyAMA1", 9600)
		for l := range latlonchannel {
			broadcast(l.String())
			fmt.Println(l)
		}
	}()

	http.HandleFunc("/gps", gps)
	log.Error().Err(http.ListenAndServe("localhost:6165", nil))
}
