package main

import (
	"context"
	"log"
	"s21_go/internal/client"
	"s21_go/internal/client/proto"

	"google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("could not connect: %v", err)
    }
    defer conn.Close()


    c := client.NewTransmitterClient(conn)


    data := &proto.TransmitterData{

    }
    if err := c.TransmitData(context.Background(), data); err != nil {
        log.Fatalf("error while calling TransmitData: %v", err)
    }
}
