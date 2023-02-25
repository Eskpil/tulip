package discovery

import (
	"bytes"
	"net"
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

func (s *Server) Handle(request *Packet, addr *net.UDPAddr) error {
	s.Transactions[request.Id] = TransactionStateRequested

	response := new(Packet)
	response.Response = new(ResponseBody)

	response.Id = request.Id
	response.Op = OpResponse

	response.Response.Address = "192.168.0.38"
	response.Response.Port = 8765
	response.Response.Version = 12
	response.Response.PublicKey = "123"
	response.Response.PrivateKey = "1234"

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

				s.Handle(packet, addr)
			}

		}
	}()

	return stopChannel, nil
}
