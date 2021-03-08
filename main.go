package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Defino una variable upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func homePage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "This is home page!")
	http.ServeFile(w, r, "./index.html")
}

// func reader
func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Cliente conectado satisfactoriamente ... ")

	reader(ws)

}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/wsocket", wsEndpoint)
}

func main() {
	fmt.Println("WebSockets!")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
