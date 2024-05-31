package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"
)


type Client struct {
	conn *grpc.ClientConn
}


func NewClient(address string) (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn}, nil
}


func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) ExampleRPC(ctx context.Context) error {
	log.Println("Example gRPC call")
	return nil
}
