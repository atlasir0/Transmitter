package client

import (
	"context"
	"math"

	"github.com/sirupsen/logrus"

	"s21_go/internal/client/proto"

	"google.golang.org/grpc"
)

type Client struct {
	client proto.TransmitterClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		client: proto.NewTransmitterClient(conn),
	}
}

func (c *Client) SendData(ctx context.Context, data *proto.TransmitterData) error {
	stream, err := c.client.TransmitData(ctx)
	if err != nil {
		logrus.Errorf("Error opening stream: %v", err)
		return err
	}
	defer stream.CloseSend()

	var (
		count          int
		totalFrequency float64
		totalSqDiff    float64
		mean           float64
		stdDev         float64
	)

	for {
		response, err := stream.Recv()
		if err != nil {
			logrus.Errorf("Error receiving data: %v", err)
			return err
		}

		logrus.Infof("Received data: SessionID=%s, Frequency=%f, Timestamp=%v", response.SessionId, response.Frequency, response.Timestamp)

		count++
		frequency := response.Frequency
		totalFrequency += frequency
		mean = totalFrequency / float64(count)
		totalSqDiff += math.Pow(frequency-mean, 2)
		stdDev = math.Sqrt(totalSqDiff / float64(count))

		logrus.Infof("Current mean: %f, Standard deviation: %f", mean, stdDev)

		k := 2.0
		if math.Abs(frequency-mean) > k*stdDev {
			logrus.Warnf("Anomaly detected: Frequency=%f is beyond expected range", frequency)
		}
	}
}
