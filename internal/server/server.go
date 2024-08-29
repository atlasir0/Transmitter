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


package main

import (
	"fmt"
)

// Определение структуры для хранения значений
type Resp struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

// Метод для получения значений из мапы (замените на ваш метод)
func (srv *Service) GetValue() map[string]int {
	// Пример значений из мапы
	return map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
}

// Функция Param для преобразования мапы в список структур Resp
func (srv *Service) Param() []Resp {
	// Получаем значения из мапы
	valuesMap := srv.GetValue()
	// Создаем срез для хранения структур Resp
	respList := make([]Resp, 0, len(valuesMap))
	// Преобразуем мапу в список структур Resp
	for k, v := range valuesMap {
		respList = append(respList, Resp{ID: v, Value: k})
	}
	return respList
}

// Структура для сервиса
type Service struct {
	// Добавьте поля, необходимые для вашего сервиса
	model Model
}

// Структура для модели (замените на вашу реализацию)
type Model struct {
	// Добавьте поля, необходимые для вашей модели
}

func main() {
	// Создаем экземпляр сервиса
	srv := &Service{}

	// Вызываем функцию Param и выводим результат
	respList := srv.Param()
	fmt.Println(respList)
}
