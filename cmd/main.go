package main

import (
	"app/db"
	s "app/server"
	"log"
	"net/http"
)

func main() {
	server := &s.PlayerServer{Store: db.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
