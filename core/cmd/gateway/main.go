package main

import (
	"github.com/eskpil/tulip/core/internal/gateway"
	log "github.com/sirupsen/logrus"
)

func main() {
	server, err := gateway.NewServer()
	if err != nil {
		log.Fatalf("Could not create gateway server: %v", err)
	}

	server.Listen()
}
