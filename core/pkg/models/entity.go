package models

import "time"

type EntityKind string
type Topic string

type Driver string

const (
	EntityKindLight     EntityKind = "light"
	EntityKindSensor    EntityKind = "sensor"
	EntityKindSpeaker   EntityKind = "speaker"
	EntityKindInterface EntityKind = "interface"
)

const (
	DriverMQTT       Driver = "mqtt"
	DriverChromeCast Driver = "chromecast"
)

type MQTTMetadata struct {
	StateTopic   Topic `json:"state_topic"`
	CommandTopic Topic `json:"command_topic"`

	UniqueId string `json:"unique_id"`
	DeviceId string `json:"device_id"`
}

type LightMetadata struct {
	Clrm                bool     `json:"clrm"`
	SupportedColorModes []string `json:"supported_color_modes"`
}

type SensorMetadata struct {
	DeviceClass       string `json:"device_class"`
	UnitOfMeasurement string `json:"unit_of_measurement"`
	StatisticsClass   string `json:"statistics_class"`
}

type SensorState struct {
	Value string `json:"value"`
}

type EntityState struct {
	Id       string `json:"-" gorm:"primaryKey"`
	EntityId string `json:"entity_id"`

	State string `json:"state"`
	// Attributes should be a map[string]string, but we marshal it
	// before we put it into the entity.
	Attributes string `json:"attributes"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Entity struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Driver Driver `json:"driver"`

	// What device owns us? A device id of "-" means we are
	// self-sufficient. For example a chromecast device.
	DeviceId string `json:"device_id"`

	// JSON marshaled metadata
	EntityMetadata string `json:"entity_metadata"`
	DriverMetadata string `json:"driver_metadata"`

	Name string `json:"name"`

	Kind EntityKind `json:"kind"`

	History []EntityState `json:"history" gorm:"foreignKey:EntityId"`
}

// TableName Change name of EntityState table to history
func (EntityState) TableName() string {
	return "history"
}
