package sockets

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.org/akrck02/nightfall/engine"
)

var upgrader = websocket.Upgrader{}
var connections = []*websocket.Conn{}

func Start() {

	go func() {
		http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {

			// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
			var err error
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				log.Print("upgrade failed: ", err)
				return
			}

			connections = append(connections, conn)
		})

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "websockets.html")
		})

		http.ListenAndServe("0.0.0.0:4321", nil)
	}()
}

func CloseConnections() {
	for _, conn := range connections {
		conn.Close()
	}

	connections = []*websocket.Conn{}
}

func SendFrame() {

	println("FRAME STATE:")
	for _, conn := range connections {

		println("Sending frame to ")
		println(conn)

		frame, err := engine.GetFrame()
		if nil != err {
			log.Printf("Error %s", err.Error())
		}

		err = conn.WriteMessage(websocket.BinaryMessage, frame)
		if nil != err {
			log.Printf("Error %s", err.Error())
		}
	}
}
