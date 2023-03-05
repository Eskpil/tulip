package main

import (
	"github.com/mochi-co/mqtt/v2"
	"github.com/mochi-co/mqtt/v2/hooks/auth"
	"github.com/mochi-co/mqtt/v2/listeners"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	server := mqtt.New(nil)
	if err := server.AddHook(new(auth.AllowHook), nil); err != nil {
		log.Errorf("Failed to add auth.AllowHook: %v", err)
	}

	tcpListener := listeners.NewTCP("tcp_1", "0.0.0.0:1883", nil)
	if err := server.AddListener(tcpListener); err != nil {
		log.Fatalf("Failed to initialize a tcp listener: %v", err)
	}

	wsListener := listeners.NewWebsocket("ws_1", "0.0.0.0:1882", nil)
	if err := server.AddListener(wsListener); err != nil {
		log.Fatalf("Failed to initialize a websocket listener: %v", err)
	}

	statsListener := listeners.NewHTTPStats("stats_1", "0.0.0.0:1881", nil, server.Info)
	if err := server.AddListener(statsListener); err != nil {
		log.Fatalf("Failed to initialize a http stats listener: %v", err)
	}

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	<-done
	log.Infof("Stopping")
}
