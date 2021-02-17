package main

import (
	"flag"
	"github.com/CoderYihaoWang/gomoku/internal/server"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()

	s := server.New()
	go s.Run()

	http.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
		server.Serve(s, w, r)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}
