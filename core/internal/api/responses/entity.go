package responses

import "github.com/eskpil/tulip/core/pkg/models"

type Entity struct {
	ID     string        `json:"id"`
	Driver models.Driver `json:"driver"`

	DeviceId string `json:"device_id"`

	EntityMetadata map[string]interface{} `json:"entity_metadata"`
	DriverMetadata map[string]interface{} `json:"driver_metadata"`

	Name string            `json:"name"`
	Kind models.EntityKind `json:"kind"`

	History []*State `json:"history"`
}
