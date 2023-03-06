package main

import (
	"context"
	"github.com/eskpil/tulip/core/internal/ca"
	"github.com/eskpil/tulip/core/pkg/pki"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type PkiServer struct {
	pki.UnimplementedPkiServer
}

func (server *PkiServer) RequestSignedCertificate(ctx context.Context, request *pki.RequestSignedCertificateRequest) (*pki.RequestSignedCertificateResponse, error) {
	cert, key, err := ca.CreatePair(request.Entity)
	if err != nil {
		return nil, err
	}

	response := new(pki.RequestSignedCertificateResponse)

	response.PrivateKey = key
	response.PublicKey = cert

	return response, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8006")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pki.RegisterPkiServer(s, &PkiServer{})

	log.Infof("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
