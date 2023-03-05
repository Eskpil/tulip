package main

import (
	"context"
	"github.com/eskpil/tulip/core/internal/api/responses"
	"github.com/eskpil/tulip/core/internal/database"
	"github.com/eskpil/tulip/core/pkg/gateway"
	"github.com/eskpil/tulip/core/pkg/models"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/eskpil/tulip/core/pkg/api"

	log "github.com/sirupsen/logrus"

	routeHandlers "github.com/eskpil/tulip/core/internal/api"
)

type ApiServer struct {
	api.UnimplementedApiServer
}

func (a ApiServer) EntityExists(ctx context.Context, request *api.EntityExistsRequest) (*api.EntityExistsResponse, error) {
	var result struct {
		Found bool
	}

	r := database.Client().Raw("SELECT EXISTS(SELECT 1 FROM entities WHERE id = ?) AS found",
		request.Id).Scan(&result)

	if r.Error != nil {
		log.Errorf("Could not check if record exists: %v", r.Error)
	}

	response := new(api.EntityExistsResponse)
	response.Ok = result.Found

	return response, nil
}

func (a ApiServer) CreateEntity(ctx context.Context, request *api.GenericEntityRequest) (*api.CreateEntityResponse, error) {
	entity := new(models.Entity)

	entity.ID = request.GetId()
	entity.Driver = models.Driver(request.GetDriver())

	entity.DeviceId = request.GetDeviceId()

	entity.EntityMetadata = string(request.GetEntityMetadata())
	entity.DriverMetadata = string(request.GetDriverMetadata())

	entity.Name = request.GetName()
	entity.Kind = models.EntityKind(request.GetKind())

	result := database.Client().Create(entity)
	if result.Error != nil {
		log.Errorf("Could not create enitity: %v", result.Error)
		return nil, result.Error
	}

	response := new(api.CreateEntityResponse)
	response.Ok = true

	return response, nil
}

func (a ApiServer) UpdateEntity(ctx context.Context, request *api.GenericEntityRequest) (*api.UpdateEntityResponse, error) {
	entity := new(models.Entity)

	entity.ID = request.GetId()
	entity.Driver = models.Driver(request.GetDriver())

	entity.DeviceId = request.GetDeviceId()

	entity.EntityMetadata = string(request.GetEntityMetadata())
	entity.DriverMetadata = string(request.GetDriverMetadata())

	entity.Name = request.GetName()
	entity.Kind = models.EntityKind(request.GetKind())

	result := database.Client().Save(entity)
	if result.Error != nil {
		log.Errorf("Could not create enitity: %v", result.Error)
		return nil, result.Error
	}

	response := new(api.UpdateEntityResponse)
	response.Ok = true

	return response, nil
}

func (a ApiServer) AppendEntityHistory(ctx context.Context, request *api.AppendEntityHistoryRequest) (*api.AppendEntityHistoryResponse, error) {
	state := new(models.EntityState)

	state.Id = uuid.New().String()
	state.State = request.GetState()
	state.EntityId = request.GetEntityId()

	result := database.Client().Create(state)
	if result.Error != nil {
		log.Errorf("Could not create state: %v", result.Error)
		return nil, result.Error
	}

	grpcConn, err := grpc.Dial("localhost:8005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := gateway.NewGatewayClient(grpcConn)

	if err != nil {
		return nil, err
	}

	responseState, err := responses.FromEntityState(*state)
	if err != nil {
		return nil, err
	}

	stateBytes, err := responseState.Marshal()
	if err != nil {
		return nil, err
	}

	gatewayRequest := new(gateway.PublishRequest)
	gatewayRequest.Subject = "state"
	gatewayRequest.Payload = stateBytes

	_, err = client.Publish(ctx, gatewayRequest)
	if err != nil {
		return nil, err
	}

	response := new(api.AppendEntityHistoryResponse)
	response.Ok = true

	return response, nil
}

func main() {
	db := database.Initialize()
	_ = db

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(g *sync.WaitGroup) {
		e := echo.New()

		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowHeaders:     []string{"*"},
			AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodDelete, http.MethodPost, http.MethodPut},
			AllowCredentials: true,
		}))
		e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))

		e.GET("/entities/", routeHandlers.GetAll)

		e.PATCH("/entities/:id/action/", routeHandlers.EntityAction)
		e.GET("/entities/:id/history/last/", routeHandlers.LastState)

		e.Logger.Fatal(e.Start(":8000"))

		g.Done()
	}(wg)

	wg.Add(1)
	go func(g *sync.WaitGroup) {
		listener, err := net.Listen("tcp", ":8001")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		api.RegisterApiServer(s, &ApiServer{})

		log.Infof("grpc server listening at %v", listener.Addr())
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

		g.Done()
	}(wg)

	wg.Wait()
}
