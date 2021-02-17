package main

import (
	"flag"
	server2 "github.com/CoderYihaoWang/gomoku/internal/server"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()

	server := server2.NewServer()
	go server.Run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server2.Serve(server, w, r)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}
