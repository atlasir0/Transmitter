package server

import (
	"io"

	"s21_go/internal/server/proto"
	"s21_go/pkg/db"

	"github.com/sirupsen/logrus"
)

func (s *TransmitterServer) TransmitData(stream proto.Transmitter_TransmitDataServer) error {
	for {
		transmission, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				logrus.Info("Client finished sending data")
				return nil
			}
			logrus.Errorf("Error receiving data: %v", err)
			return err
		}

		transmitterData := &db.TransmitterData{
			SessionID: transmission.SessionId,
			Frequency: transmission.Frequency,
			Timestamp: transmission.Timestamp.AsTime(),
		}
		logrus.Infof("Received data: SessionID=%s, Frequency=%f, Timestamp=%v", transmission.SessionId, transmission.Frequency, transmission.Timestamp.AsTime())

		orm := db.NewORM(s.DB.DB)
		err = orm.CreateTransmission(transmitterData)
		if err != nil {
			logrus.Errorf("Error saving data to the database: %v", err)
			return err
		}
		logrus.Info("Data saved to the database successfully")
	}
}
