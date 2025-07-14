package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// WebSocket handling

	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Received Message: %s\n", data)

		var updatedData string = "Received Data : " + string(data)
		//Echo
		err = conn.WriteMessage(messageType, []byte(updatedData))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func main() {
	http.HandleFunc("/ws", handleWebsocket)
	fmt.Println("Websocket Server is running on :80/ws")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}
}
