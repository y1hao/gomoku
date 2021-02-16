package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()

	wsServer := NewWsServer()
	go wsServer.Run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(wsServer, w, r)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}
