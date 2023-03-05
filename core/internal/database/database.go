package database

import (
	"github.com/eskpil/tulip/core/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

var conn *gorm.DB

func Initialize() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect with database: %v", err)
	}

	// Migrate the schema
	_ = db.AutoMigrate(&models.Device{})
	_ = db.AutoMigrate(&models.Entity{})
	_ = db.AutoMigrate(&models.EntityState{})

	conn = db

	return db
}

func Client() *gorm.DB {
	if conn == nil {
		log.Fatalf("Database has not been initialized yet, please call Initialize")
	}

	return conn
}
