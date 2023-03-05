package gateway

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"net"
	"sync"

	gatewayApi "github.com/eskpil/tulip/core/pkg/gateway"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	gatewayApi.UnimplementedGatewayServer

	E *echo.Echo

	Sockets  map[string]*websocket.Conn
	Upgrader websocket.Upgrader

	ApiListener net.Listener
	Api         *grpc.Server
}

func NewServer() (*Server, error) {
	server := new(Server)

	server.E = echo.New()
	server.E.GET("/gateway/", handleConnection(server))

	server.Upgrader = websocket.Upgrader{}
	server.Sockets = make(map[string]*websocket.Conn)

	listener, err := net.Listen("tcp", ":8005")
	if err != nil {
		return nil, err
	}

	{
		s := grpc.NewServer()
		gatewayApi.RegisterGatewayServer(s, server)
		server.ApiListener = listener
		server.Api = s
	}

	return server, nil
}

func (s *Server) Publish(ctx context.Context, request *gatewayApi.PublishRequest) (*gatewayApi.PublishResponse, error) {
	message := new(Message)

	message.Subject = request.GetSubject()
	message.EntityId = request.GetEntityId()

	if err := json.Unmarshal(request.GetPayload(), &message.Payload); err != nil {
		return nil, err
	}

	for id, socket := range s.Sockets {
		if socket == nil {
			continue
		}

		if err := socket.WriteJSON(message); err != nil {
			if err != nil {
				s.Sockets[id] = nil
			}
		}
	}

	response := new(gatewayApi.PublishResponse)
	response.Ok = true

	return response, nil
}

func (s *Server) Listen() {

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(g *sync.WaitGroup) {

		s.E.Logger.Fatal(s.E.Start(":8004"))

		g.Done()
	}(wg)

	wg.Add(1)
	go func(g *sync.WaitGroup) {
		log.Infof("grpc server listening at %v", s.ApiListener.Addr())
		if err := s.Api.Serve(s.ApiListener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

		g.Done()
	}(wg)

	wg.Wait()
}
