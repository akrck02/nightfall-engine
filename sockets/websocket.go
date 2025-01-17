package sockets

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.org/akrck02/nightfall/constants"
)

var upgrader = websocket.Upgrader{}
var connections = []*websocket.Conn{} 

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

func CloseConnections() {
  for _,conn := range connections {
    conn.Close()
  }

  connections = []*websocket.Conn{}
}

func SendFrameAsHtml(frame [][]uint32) {

  var html = ``
  for i, pixel := range frame {
    if i % constants.Resolution[1] == 0 {
      html += `</div><div id="display-row" style="display:flex;flex-direction:row;">`
    } 
    
    html += fmt.Sprintf(`<div id="pixel" style="background:rgba(%d,%d,%d,%d);width:1rem;height:1rem;color:transparent;">.</div>`, pixel[0],pixel[1],pixel[2], 1)
  }

  println("FRAME STATE:")
  for _, conn := range connections {

    println("Sending frame to ")
    println(conn)
    
    err := conn.WriteMessage(websocket.BinaryMessage, []byte(html))
    if (nil != err){
      print("ERROR: ")
      println(err.Error())
    }
  }
}
