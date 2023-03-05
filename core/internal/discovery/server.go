package discovery

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/eskpil/tulip/core/pkg/wind"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	Wind *wind.Server
	Mqtt mqtt.Client
}

func NewServer() (*Server, error) {
	server := new(Server)

	windServer, err := wind.NewServer()
	if err != nil {
		return nil, fmt.Errorf("wind: %v", err)
	}
	server.Wind = windServer

	mqtt.ERROR = log.New()
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("tulip_core")

	opts.SetKeepAlive(60 * time.Second)
	// Set the message callback handler
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	server.Mqtt = mqtt.NewClient(opts)
	if token := server.Mqtt.Connect(); token.Wait() && token.Error() != nil {
		return nil, err
	}

	return server, nil
}

func (s *Server) Listen() (chan bool, error) {
	stopChannel := make(chan bool)

	if err := s.Wind.Listen(stopChannel); err != nil {
		return stopChannel, err
	}

	// Subscribe to a topic
	go func() {
		if token := s.Mqtt.Subscribe("#", 0, nil); token.Wait() && token.Error() != nil {
			log.Errorf("Failed to subscribe: %v", token.Error())
		}

		<-stopChannel

		// Unsubscribe
		if token := s.Mqtt.Unsubscribe("#"); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}

		// Disconnect
		s.Mqtt.Disconnect(250)
	}()

	return stopChannel, nil
}
