package utils

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/eskpil/tulip/core/pkg/api"
	"github.com/eskpil/tulip/core/pkg/models"
)

func PublishState(ctx context.Context, state *models.EntityState) (bool, error) {
	grpcConn, err := grpc.Dial("localhost:8003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := api.NewApiClient(grpcConn)

	if err != nil {
		return false, err
	}

	body := new(api.AppendEntityHistoryRequest)
	body.EntityId = state.EntityId
	body.State = state.State

	response, err := client.AppendEntityHistory(ctx, body)
	if err != nil {
		return false, err
	}

	return response.Ok, err
}
