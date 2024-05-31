package server

import (
	"net"

	"s21_go/internal/server/proto"
	"s21_go/pkg/config"
	"s21_go/pkg/db"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type TransmitterServer struct {
	proto.UnimplementedTransmitterServer
	DB *db.DB
}

func StartServer() error {
	logrus.Info("Starting server...")

	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("Failed to load config: %v", err)
		return err
	}
	logrus.Info("Config loaded successfully")

	database, err := db.ConnectDB(cfg)
	if err != nil {
		logrus.Fatalf("Failed to connect to the database: %v", err)
		return err
	}
	logrus.Info("Connected to the database successfully")

	grpcServer := grpc.NewServer()
	transmitterServer := &TransmitterServer{
		DB: database,
	}
	proto.RegisterTransmitterServer(grpcServer, transmitterServer)
	logrus.Info("gRPC server and TransmitterServer registered successfully")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logrus.Fatalf("Failed to listen on port 50051: %v", err)
		return err
	}
	logrus.Info("Server is listening on port 50051")

	err = grpcServer.Serve(listener)
	if err != nil {
		logrus.Fatalf("Failed to serve: %v", err)
		return err
	}
	logrus.Info("Server started successfully")
	return nil
}
