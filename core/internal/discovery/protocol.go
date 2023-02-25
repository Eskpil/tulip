package discovery

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Op uint16

const (
	OpRequest Op = iota
	OpResponse
)

type RequestBody struct {
	Version  uint16
	Hostname string
	Uname    string
}

type ResponseBody struct {
	Address string

	PrivateKey string
	PublicKey  string

	Port    uint16
	Version uint16
}

type Packet struct {
	Id uint16
	Op Op

	Request  *RequestBody
	Response *ResponseBody
}

func EncodeString(buffer *bytes.Buffer, data string) error {
	if err := binary.Write(buffer, binary.LittleEndian, uint32(len(data))); err != nil {
		return err
	}

	buffer.Write([]byte(data))

	return nil
}

func Encode(packet Packet) (*bytes.Buffer, error) {
	buffer := &bytes.Buffer{}

	if err := binary.Write(buffer, binary.LittleEndian, packet.Id); err != nil {
		return nil, err
	}

	if err := binary.Write(buffer, binary.LittleEndian, packet.Op); err != nil {
		return nil, err
	}

	if packet.Op == OpRequest {
		if packet.Request == nil {
			return nil, fmt.Errorf("request is nil")
		}

		if err := binary.Write(buffer, binary.LittleEndian, packet.Request.Version); err != nil {
			return nil, err
		}
	} else {
		//if err := EncodeAddress(buffer, packet.Response.Address); err != nil {
		//	return nil, err
		//}
		if err := EncodeString(buffer, packet.Response.Address); err != nil {
			return nil, err
		}
		if err := EncodeString(buffer, packet.Response.PrivateKey); err != nil {
			return nil, err
		}
		if err := EncodeString(buffer, packet.Response.PublicKey); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.LittleEndian, packet.Response.Port); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.LittleEndian, packet.Response.Version); err != nil {
			return nil, err
		}
	}

	return buffer, nil
}

func DecodeString(buffer *bytes.Buffer) (string, error) {
	var size uint32
	if err := binary.Read(buffer, binary.LittleEndian, &size); err != nil {
		return "", err
	}

	buf := make([]byte, size)
	if _, err := buffer.Read(buf); err != nil {
		return "", err
	}

	return string(buf), nil
}

func Decode(buffer *bytes.Buffer) (*Packet, error) {
	packet := new(Packet)

	if err := binary.Read(buffer, binary.LittleEndian, &packet.Id); err != nil {
		return nil, err
	}

	if err := binary.Read(buffer, binary.LittleEndian, &packet.Op); err != nil {
		return nil, err
	}

	if packet.Op == OpRequest {
		packet.Request = new(RequestBody)

		if err := binary.Read(buffer, binary.LittleEndian, &packet.Request.Version); err != nil {
			return nil, err
		}

		if hostname, err := DecodeString(buffer); err != nil {
			return nil, err
		} else {
			packet.Request.Hostname = hostname
		}

		if uname, err := DecodeString(buffer); err != nil {
			return nil, err
		} else {
			packet.Request.Uname = uname
		}
	}

	return packet, nil
}
