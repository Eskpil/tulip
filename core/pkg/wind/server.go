package wind

import (
	"bytes"
	"context"
	"fmt"
	"github.com/eskpil/tulip/core/pkg/api"
	"github.com/eskpil/tulip/core/pkg/models"
	"github.com/eskpil/tulip/core/pkg/pki"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"time"
)
import log "github.com/sirupsen/logrus"

type Server struct {
	Conn *net.UDPConn

	Transactions map[uint16]TransactionState
}

func genericFromEntity(entity *models.Entity) *api.GenericEntityRequest {
	generic := new(api.GenericEntityRequest)

	generic.Id = entity.ID
	generic.Driver = string(entity.Driver)

	generic.DeviceId = entity.DeviceId

	generic.EntityMetadata = []byte(entity.EntityMetadata)
	generic.DriverMetadata = []byte(entity.DriverMetadata)

	generic.Name = entity.Name
	generic.Kind = string(entity.Kind)

	return generic
}

func NewServer() (*Server, error) {
	server := new(Server)

	server.Transactions = make(map[uint16]TransactionState)

	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:7654")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("could not listen: (%v)", err)
	}

	server.Conn = conn

	return server, nil
}

func (s *Server) RequestCertificates(ctx context.Context, entity string) (*pki.RequestSignedCertificateResponse, error) {
	grpcConn, err := grpc.Dial("localhost:8006", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pki.NewPkiClient(grpcConn)

	body := new(pki.RequestSignedCertificateRequest)
	body.Entity = entity

	certificates, err := client.RequestSignedCertificate(ctx, body)
	if err != nil {
		return nil, err
	}

	return certificates, nil
}

func (s *Server) CreateOrUpdateEntity(ctx context.Context, request *Packet) (bool, error) {
	grpcConn, err := grpc.Dial("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return false, err
	}

	client := api.NewApiClient(grpcConn)

	entity := new(models.Entity)

	entity.ID = fmt.Sprintf("%s.%s", request.Request.DeviceClass, request.Request.Hostname)
	entity.Name = request.Request.Hostname
	entity.Kind = models.EntityKindInterface
	entity.Driver = models.DriverWind

	entityExistsBody := new(api.EntityExistsRequest)
	entityExistsBody.Id = entity.ID

	entityExistsResponse, err := client.EntityExists(ctx, entityExistsBody)
	if err != nil {
		return false, err
	}

	if entityExistsResponse.Ok {
		updateEntityBody := genericFromEntity(entity)

		updateEntityResponse, err := client.UpdateEntity(ctx, updateEntityBody)
		if err != nil {
			return false, err
		}

		return updateEntityResponse.Ok, nil
	} else {
		createEntityBody := genericFromEntity(entity)

		createEntityResponse, err := client.CreateEntity(ctx, createEntityBody)
		if err != nil {
			return false, err
		}

		return createEntityResponse.Ok, nil
	}
}

func (s *Server) Handle(request *Packet, addr *net.UDPAddr) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.Transactions[request.Id] = TransactionStateRequested

	response := new(Packet)
	response.Response = new(ResponseBody)

	response.Id = request.Id
	response.Op = OpResponse

	certificates, err := s.RequestCertificates(ctx, fmt.Sprintf("%s.%s", request.Request.DeviceClass, request.Request.Hostname))
	if err != nil {
		return err
	}

	response.Response.Address = "192.168.0.38"
	response.Response.Version = 12
	response.Response.PublicKey = certificates.GetPublicKey()
	response.Response.PrivateKey = certificates.GetPrivateKey()

	httpSupportedProtocol := SupportedProtocol{Protocol: "http", Port: 8000}
	apiService := Service{Name: "api", SupportedProtocols: []SupportedProtocol{httpSupportedProtocol}}
	response.Response.Services = append(response.Response.Services, apiService)

	log.Infof("Sending response")

	encoded, err := Encode(*response)
	if err != nil {
		return err
	}

	addr.Port = 6543

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}

	if _, err = conn.Write(encoded.Bytes()); err != nil {
		return err
	}

	if _, err := s.CreateOrUpdateEntity(ctx, request); err != nil {
		return err
	}

	return nil
}

func (s *Server) Listen(stop chan bool) error {

	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				buf := make([]byte, 1024)

				_, addr, err := s.Conn.ReadFromUDP(buf)
				if err != nil {
					log.Errorf("could not read: (%v)", err)
					continue
				}

				buffer := bytes.NewBuffer(buf)

				packet, err := Decode(buffer)
				if err != nil {
					log.Errorf("could not decode: (%v)", err)
					continue
				}

				if err := s.Handle(packet, addr); err != nil {
					log.Errorf("Failed to handle: (%v)", err)
				}
			}

		}
	}()

	return nil
}
