# Transmitter

 Проект представляет собой реализацию передатчика данных с использованием gRPC, а также логику обнаружения аномалий в потоке данных и сохранение их в базу данных PostgreSQL с использованием ORM.

## Структура проекта

project/
├── cmd/
│ ├── client/
│ │ ├── main.go
│ │ └── config.yaml
│ └── server/
│ ├── main.go
│ └── config.yaml
├── internal/
│ ├── client/
│ │ ├── client.go
│ │ ├── handler.go
│ │ └── proto/
│ │ ├── transmitter.proto
│ │ ├── transmitter.pb.go
│ │ └── transmitter_grpc.pb.go
│ └── server/
│ ├── server.go
│ ├── handler.go
│ └── proto/
│ ├── transmitter.proto
│ ├── transmitter.pb.go
│ └── transmitter_grpc.pb.go
└── pkg/
├── config/
│ └── config.go
├── db/
│ ├── db.go
│ ├── models.go
│ └── orm.go
└── migration/
└── migration.go


## Использование

### Сервер

1. Скомпилировать сервер с помощью команды `make server`.
2. Запустить сервер  `go run  bin/server`

### Клиент

1. Скомпилировать клиент с помощью команды `make client`.
2. Запустить клиента `go run  bin/client` 

