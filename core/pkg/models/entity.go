package models

type EntityKind string
type Topic string

const (
	EntityKindLight  EntityKind = "light"
	EntityKindSensor EntityKind = "sensor"
)

type Entity struct {
	ID       string `json:"id" gorm:"primaryKey"`
	DeviceId string `json:"device_id"`

	Kind     EntityKind `json:"kind"`
	Metadata string     `json:"metadata"`

	StateTopic   Topic `json:"state_topic"`
	CommandTopic Topic `json:"command_topic"`

	UniqueId string `json:"unique_id"`
}
