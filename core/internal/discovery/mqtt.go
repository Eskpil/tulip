package discovery

import (
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/eskpil/tulip/core/pkg/api"
	"github.com/eskpil/tulip/core/pkg/models"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type Device struct {
	Ids          string `json:"ids"`
	Name         string `json:"name"`
	Software     string `json:"sw"`
	Model        string `json:"model"`
	Manufacturer string `json:"mf"`
}

type Config struct {
	Schema string `json:"schema"`

	DevClass          string `json:"dev_cla"`
	UnitOfMeasurement string `json:"unit_of_meas"`
	StatisticsClass   string `json:"stat_cla"`

	Clrm                bool     `json:"clrm"`
	SupportedColorModes []string `json:"supported_color_modes"`

	Name string `json:"name"`

	StateTopic        string `json:"stat_t"`
	CommandTopic      string `json:"cmd_t"`
	AvailabilityTopic string `json:"avty_y"`

	UniqueId string `json:"uniq_id"`

	Dev Device `json:"dev"`
}

func createClient() (api.ApiClient, error) {
	grpcConn, err := grpc.Dial("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := api.NewApiClient(grpcConn)

	if err != nil {
		return client, err
	}

	return client, err
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

func entityExists(ctx context.Context, id string) (bool, error) {
	client, err := createClient()
	if err != nil {
		return false, err
	}

	body := new(api.EntityExistsRequest)
	body.Id = id

	response, err := client.EntityExists(ctx, body)
	if err != nil {
		return false, err
	}

	return response.Ok, nil
}

func createEntity(ctx context.Context, entity *models.Entity) (bool, error) {
	client, err := createClient()
	if err != nil {
		return false, err
	}

	body := genericFromEntity(entity)

	response, err := client.CreateEntity(ctx, body)
	if err != nil {
		return false, err
	}

	return response.Ok, nil
}

func updateEntity(ctx context.Context, entity *models.Entity) (bool, error) {
	client, err := createClient()
	if err != nil {
		return false, err
	}

	body := genericFromEntity(entity)

	response, err := client.UpdateEntity(ctx, body)
	if err != nil {
		return false, err
	}

	return response.Ok, nil
}

func appendEntityHistory(ctx context.Context, state *models.EntityState) (bool, error) {
	client, err := createClient()
	if err != nil {
		return false, nil
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

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	parts := strings.Split(msg.Topic(), "/")
	action := parts[len(parts)-1]

	log.Infof("topic: (%s)", msg.Topic())

	if action == "config" {
		var config Config
		if err := json.Unmarshal(msg.Payload(), &config); err != nil {
			log.Errorf("failed to unmarshal config: %v", err)
		}

		mqttMetadata := new(models.MQTTMetadata)

		mqttMetadata.CommandTopic = models.Topic(config.CommandTopic)
		mqttMetadata.StateTopic = models.Topic(config.StateTopic)
		mqttMetadata.UniqueId = config.UniqueId

		mqttMetadataBytes, err := json.Marshal(mqttMetadata)
		if err != nil {
			log.Errorf("Could not marshal mqtt metadata: %v", err)
			return
		}

		entity := new(models.Entity)
		entityId := fmt.Sprintf("%s.%s", parts[1], parts[3])

		entityExists, err := entityExists(ctx, entityId)
		if err != nil {
			log.Errorf("Could not determine if entity already exists: %v", err)
			return
		}

		entity.ID = entityId
		entity.Driver = models.DriverMQTT
		entity.DriverMetadata = string(mqttMetadataBytes)

		entity.Name = config.Name

		switch parts[1] {
		case "light":
			{
				entity.Kind = models.EntityKindLight
				lightMetadata := new(models.LightMetadata)

				lightMetadata.Clrm = config.Clrm
				lightMetadata.SupportedColorModes = config.SupportedColorModes

				lightMetadataBytes, err := json.Marshal(lightMetadata)
				if err != nil {
					log.Errorf("Could not marshal lightMetadata: %v", err)
					return
				}
				entity.EntityMetadata = string(lightMetadataBytes)
			}
		case "sensor":
			{
				entity.Kind = models.EntityKindSensor
				sensorMetadata := new(models.SensorMetadata)

				sensorMetadata.UnitOfMeasurement = config.UnitOfMeasurement
				sensorMetadata.DeviceClass = config.DevClass
				sensorMetadata.StatisticsClass = config.StatisticsClass

				sensorMetadataBytes, err := json.Marshal(sensorMetadata)
				if err != nil {
					log.Errorf("Could not marshal sensor metadata: %v", err)
					return
				}
				entity.EntityMetadata = string(sensorMetadataBytes)
			}
		default:
			{
				log.Errorf("Unhandled mqtt kind: %s", parts[1])
				return
			}
		}

		// Either insert or update the entity in the database.
		if !entityExists {
			if _, err := createEntity(ctx, entity); err != nil {
				log.Errorf("Failed to create entity: %v", err)
			}
		} else {
			if _, err := updateEntity(ctx, entity); err != nil {
				log.Errorf("Failed to update entity: %v", err)
			}
		}
	} else if action == "state" {
		log.Infof("Topic: (%v)", msg.Topic())

		state := new(models.EntityState)
		state.Id = uuid.New().String()
		state.EntityId = fmt.Sprintf("%s.%s", parts[1], parts[2])
		state.State = string(msg.Payload())

		if _, err := appendEntityHistory(ctx, state); err != nil {
			log.Errorf("Could not add entity history: %v", err)
			return
		}
	} else {
	}
}
