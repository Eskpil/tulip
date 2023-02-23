package main

import (
	"github.com/eskpil/tulip/core/pkg/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Device{})
	db.AutoMigrate(&models.Entity{})

	log.Info("hello, world")
}
