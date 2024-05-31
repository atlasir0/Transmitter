package db

import (
	"fmt"
	"log"
	"s21_go/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func ConnectDB(cfg *config.Config) (*DB, error) {
	dsn := fmt.Sprintf("user=atlasiro password=123 dbname=db host=localhost port=5432 sslmode=disable")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&TransmitterData{})
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database and migrated the schema")
	return &DB{db}, nil
}
