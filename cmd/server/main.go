package main

import (
	"s21_go/internal/server"
	"s21_go/pkg/config"
	"s21_go/pkg/db"
	"s21_go/pkg/migration"

	"github.com/sirupsen/logrus"
)

func Start(cfg *config.Config) error {
	logrus.Info("Starting application...")

	database, err := db.ConnectDB(cfg)
	if err != nil {
		logrus.Fatalf("Failed to connect to the database: %v", err)
		return err
	}
	logrus.Info("Connected to the database successfully")

	err = migration.MigrateDB(database.DB)
	if err != nil {
		logrus.Fatalf("Failed to migrate the database: %v", err)
		return err
	}
	logrus.Info("Database migration completed successfully")

	err = server.StartServer()
	if err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
		return err
	}
	logrus.Info("Server started successfully")
	return nil
}

func main() {
	logrus.Info("Loading configuration...")
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("Failed to load config: %v", err)
	}

	if err := Start(cfg); err != nil {
		logrus.Fatalf("Failed to start application: %v", err)
	}
	logrus.Info("Application started successfully")
}
