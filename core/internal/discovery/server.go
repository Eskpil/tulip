package discovery

import (
	"bytes"
	"context"
	"fmt"
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
	grpcConn, err := grpc.Dial("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := pki.NewPkiClient(grpcConn)

	body := new(pki.RequestSignedCertificateRequest)
	body.Entity = entity

	certificates, err := client.RequestSignedCertificate(ctx, body)
	if err != nil {
		return nil, err
	}

	return certificates, nil
}

func (s *Server) Handle(request *Packet, addr *net.UDPAddr) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.Transactions[request.Id] = TransactionStateRequested

	response := new(Packet)
	response.Response = new(ResponseBody)

	response.Id = request.Id
	response.Op = OpResponse

	certificates, err := s.RequestCertificates(ctx, fmt.Sprintf("interface.%s", request.Request.Hostname))

	response.Response.Address = "192.168.0.38"
	response.Response.Version = 12
	response.Response.PublicKey = certificates.GetPublicKey()
	response.Response.PrivateKey = certificates.GetPrivateKey()

	deviceService := Service{Name: "device", Port: 8002}
	response.Response.Services = append(response.Response.Services, deviceService)

	entityService := Service{Name: "entities", Port: 8003}
	response.Response.Services = append(response.Response.Services, entityService)

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

	return nil
}

func (s *Server) Listen() (chan bool, error) {
	stopChannel := make(chan bool)

	go func() {
		for {
			select {
			case <-stopChannel:
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

	return stopChannel, nil
}
