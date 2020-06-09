package main

import (
	"github.com/R3l3ntl3ss/Jot/server"
	"log"
	"net/http"
)

func main() {
	r := server.NewRouter()

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Unable to start Server", err)
	}
}
