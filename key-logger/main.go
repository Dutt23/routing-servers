package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			fmt.Printf("%+v\n", r)
			return true
		},
	}

	listenAddr string
	wsAddr     string
	jsTemplate *template.Template
)

func init() {
	flag.StringVar(&listenAddr, "listen-addr", "", "Address to listen on")
	flag.StringVar(&wsAddr, "ws-addr", "", "Address for web socker to listen on")
	flag.Parse()
	var err error
	jsTemplate, err = template.ParseFiles("logger.js")
	if err != nil {
		log.Panic("err")
	}
}

func serveFile(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/javascript")
	jsTemplate.Execute(writer, &wsAddr)
}

func serveWS(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	fmt.Printf("Connection from %s\n", conn.RemoteAddr().String())
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Printf("From %s: %s\n", conn.RemoteAddr().String(), string(msg))
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ws", serveWS)
	router.HandleFunc("/k.js", serveFile)
	log.Fatal(http.ListenAndServe(":8080", router))
}
