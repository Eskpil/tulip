package main

import (
	"github.com/eskpil/tulip/core/internal/discovery"
	log "github.com/sirupsen/logrus"
)

func main() {
	server, err := discovery.NewServer()
	if err != nil {
		log.Fatalf("Failed to create a new discovery server: (%v)", err)
	}

	quit, err := server.Listen()
	if err != nil {
		log.Fatalf("Failed to listen for new entities: (%v)", err)
	}

	_ = quit

	select {}
}
