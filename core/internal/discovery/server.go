package discovery

import (
	"context"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/eskpil/tulip/core/pkg/wind"
	"google.golang.org/grpc"
	"net"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	discoveryApi "github.com/eskpil/tulip/core/pkg/discovery"
)

type Server struct {
	discoveryApi.UnimplementedDiscoveryServer

	Wind *wind.Server
	Mqtt mqtt.Client

	ApiListener net.Listener
	Api         *grpc.Server
}

func NewServer() (*Server, error) {
	server := new(Server)

	{
		windServer, err := wind.NewServer()
		if err != nil {
			return nil, fmt.Errorf("wind: %v", err)
		}
		server.Wind = windServer
	}

	{
		mqtt.ERROR = log.New()
		opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("tulip_core")

		opts.SetKeepAlive(60 * time.Second)
		// Set the message callback handler
		opts.SetDefaultPublishHandler(f)
		opts.SetPingTimeout(1 * time.Second)

		server.Mqtt = mqtt.NewClient(opts)
		if token := server.Mqtt.Connect(); token.Wait() && token.Error() != nil {
			return nil, token.Error()
		}
	}

	{
		listener, err := net.Listen("tcp", ":8003")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		discoveryApi.RegisterDiscoveryServer(s, server)

		server.ApiListener = listener
		server.Api = s

	}

	return server, nil
}

func (s *Server) PublishMQTTMessage(ctx context.Context, request *discoveryApi.PublishMQTTMessageRequest) (*discoveryApi.PublishMQTTMessageResponse, error) {
	s.Mqtt.Publish(request.GetTopic(), byte(2), true, request.GetPayload())

	response := new(discoveryApi.PublishMQTTMessageResponse)
	response.Ok = true

	return response, nil
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

	go func() {
		log.Infof("grpc server listening at %v", s.ApiListener.Addr())
		if err := s.Api.Serve(s.ApiListener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return stopChannel, nil
}
