package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()

	server := NewServer()
	go server.Run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Serve(server, w, r)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}
