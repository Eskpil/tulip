package models

import (
	"gorm.io/gorm"
	"time"
)

type Device struct {
	ID string `json:"id" gorm:"primaryKey"`

	Name         string `json:"name"`
	Software     string `json:"software"`
	Model        string `json:"model"`
	Manufacturer string `json:"manufacturer"`

	AvailabilityTopic Topic `json:"availability_topic"`

	Entities []Entity `json:"entities"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
