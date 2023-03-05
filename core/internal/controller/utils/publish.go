package utils

import (
	"context"
	"github.com/eskpil/tulip/core/pkg/discovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Publish(ctx context.Context, topic string, payload []byte) (bool, error) {
	grpcConn, err := grpc.Dial("localhost:8003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := discovery.NewDiscoveryClient(grpcConn)

	if err != nil {
		return false, err
	}

	request := new(discovery.PublishMQTTMessageRequest)
	request.Topic = topic
	request.Payload = payload

	response, err := client.PublishMQTTMessage(ctx, request)
	if err != nil {
		return false, err
	}

	return response.Ok, nil
}
