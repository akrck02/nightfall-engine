package sockets

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var connections = []*websocket.Conn{} 

// Start websocket server
func Start() {
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
}

// Close the active connections
func CloseConnections() {
  for _,conn := range connections {
    conn.Close()
  }

  connections = []*websocket.Conn{}
}

// Send a frame to every active connection
func SendFrame(frame *image.RGBA) {

  buf := new(bytes.Buffer)
  err := jpeg.Encode(buf, frame, nil)
  if nil != err {
    print("Cannot generate: frame skipped.")
    return 
  } 

  bytesToSend := buf.Bytes()

  println("FRAME STATE:")
  for _, conn := range connections {

    println(fmt.Sprintf("Sending frame to %s", conn.LocalAddr().String()))
    err := conn.WriteMessage(websocket.BinaryMessage, bytesToSend)
    if (nil != err){
      print("ERROR: ")
      println(err.Error())
    } 
  }
}
