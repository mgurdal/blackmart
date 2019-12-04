package main

import (
	"log"

	"github.com/mgurdal/blackmarkt/network"
	"github.com/mgurdal/blackmarkt/store"
)

func main() {
	store.GetStore()
	srv, err := network.NewServer(
		"tcp",
		"0.0.0.0:9001",
	)
	if err != nil {
		log.Fatal(err)
	}

	srv.Run()
}
