# Transmitter

Проект представляет собой реализацию передатчика данных с использованием gRPC, а также логику обнаружения аномалий в потоке данных и сохранение их в базу данных PostgreSQL с использованием ORM.

## Описание

### Transmitter (Передатчик)
Сервер gRPC, который принимает соединения от клиентов и передает им поток данных. Данные включают в себя уникальный идентификатор сессии (UUID), частоту (frequency) и временную метку (timestamp) в формате UTC. Частота выбирается случайным образом из нормального распределения с заданными средним значением и стандартным отклонением, которые также выбираются случайным образом для каждой новой сессии.

### Anomaly Detection (Обнаружение аномалий)
Клиентская часть, которая подключается к серверу gRPC, получает поток данных и вычисляет среднее значение и стандартное отклонение для текущей сессии. После того, как параметры распределения достаточно точно определены, клиент переходит к этапу обнаружения аномалий, используя коэффициент отклонения (k). Аномалией считается значение частоты, которое отличается от ожидаемого значения более чем на k * стандартное отклонение.

### Report (Отчет)
Все обнаруженные аномалии сохраняются в базу данных PostgreSQL с использованием ORM. Это позволяет безопасно сохранять данные без риска SQL-инъекций.

### All Together (Все вместе)
Проект объединяет сервер, клиент, логику обнаружения аномалий и сохранения данных в единую систему.

## Использование

### Сервер

1. Скомпилировать сервер с помощью команды `make server`.
2. Запустить сервер  `go run  bin/server`

### Клиент

1. Скомпилировать клиент с помощью команды `make client`.
2. Запустить клиента `go run  bin/client` 

