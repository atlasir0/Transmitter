// project/internal/client/client.go

package client

import (
	"context"
	"log"

	"s21_go/internal/client/proto"

	"google.golang.org/grpc"
)

type TransmitterClient struct {
	Conn *grpc.ClientConn
}

func NewTransmitterClient(conn *grpc.ClientConn) *TransmitterClient {
	return &TransmitterClient{
		Conn: conn,
	}
}

func (c *TransmitterClient) TransmitData(ctx context.Context, data *proto.TransmitterData) error {
	client := proto.NewTransmitterClient(c.Conn)
	stream, err := client.TransmitData(ctx)
	if err != nil {
		return err
	}
	defer stream.CloseSend()

	if err := stream.Send(data); err != nil {
		return err
	}
	response, err := stream.Recv()
	if err != nil {
		return err
	}

	log.Printf("Received response: %v", response)
	return nil
}

