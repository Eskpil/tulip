package controller

import (
	"context"
	"encoding/json"
	"github.com/eskpil/tulip/core/internal/controller/light"
	"github.com/eskpil/tulip/core/pkg/models"
	log "github.com/sirupsen/logrus"
)

func Action(ctx context.Context, entity *models.Entity, command *Command) error {
	switch command.Subject {
	case "light":
		{
			var subcommand light.Subcommand

			subcommandBytes, err := json.Marshal(command.Subcommand)
			if err != nil {
				return err
			}

			if err := json.Unmarshal(subcommandBytes, &subcommand); err != nil {
				return err
			}

			return light.Action(ctx, entity, subcommand)
		}
	default:
		log.Errorf("Unhandled subject: %s", command.Subject)
	}

	return nil
}
