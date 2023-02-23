package main

import (
	"github.com/hashicorp/mdns"
	log "github.com/sirupsen/logrus"
)

func main() {
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			log.Infof(" (%s): [%v]\n", entry.Name, entry.InfoFields)
		}
	}()

	// Start the lookup
	if err := mdns.Lookup("_googlecast._tcp", entriesCh); err != nil {
		close(entriesCh)
		log.Fatalf("Failed to lookup \"_googlecast._tcp\": %v", err)
	}

	close(entriesCh)
}
