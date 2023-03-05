package light

import (
	"context"
	"encoding/json"
	"github.com/eskpil/tulip/core/internal/controller/utils"
	"github.com/eskpil/tulip/core/internal/database"
	"github.com/eskpil/tulip/core/pkg/models"
	log "github.com/sirupsen/logrus"
)

type Subcommand struct {
	Name string `json:"name"`
}

type lightState struct {
	ColorMode string      `json:"color_mode"`
	State     string      `json:"state"`
	Color     interface{} `json:"color"`
}

type lightCommand struct {
	State string `json:"state"`
}

func Action(ctx context.Context, entity *models.Entity, command Subcommand) error {
	switch command.Name {
	case "opposite":
		{
			state := new(models.EntityState)
			result := database.Client().Order("created_at DESC").Find(state, "entity_id = ?", entity.ID)
			if result.Error != nil {
				log.Infof("Could not get the last state of entity: %s: %v", entity.ID, result.Error)
				return result.Error
			}

			log.Infof(" > state: %s", state.State)

			lState := new(lightState)
			if err := json.Unmarshal([]byte(state.State), &lState); err != nil {
				return err
			}

			switch entity.Driver {
			case models.DriverMQTT:
				{
					var metadata models.MQTTMetadata
					if err := json.Unmarshal([]byte(entity.DriverMetadata), &metadata); err != nil {
						return err
					}

					mqttCommand := new(lightCommand)
					if lState.State == "ON" {
						mqttCommand.State = "OFF"
					}
					if lState.State == "OFF" {
						mqttCommand.State = "ON"
					}

					mqttCommandBytes, err := json.Marshal(&mqttCommand)
					if err != nil {
						return err
					}

					_, err = utils.Publish(ctx, string(metadata.CommandTopic), mqttCommandBytes)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
