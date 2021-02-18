package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/CoderYihaoWang/gomoku/internal/server"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()

	s := server.New()
	go s.Run()

	http.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
		server.ServeWs(s, w, r)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}
