package models

type CardKind string

const (
	CardKindLight  CardKind = "light"
	CardKindSensor CardKind = "sensor"
	CardKindGroup  CardKind = "group"
)

type Card struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`

	Kind CardKind `json:"kind"`
	// Subkind is only relevant for groups.
	Subkind CardKind `json:"subkind"`

	Entities []Entity `json:"entities"`
}
