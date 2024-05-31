package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc"
)


type Server struct {
	address string
	server  *grpc.Server
}


func NewServer(address string) *Server {
	return &Server{address: address}
}


func (s *Server) RegisterService(service grpc.ServiceRegistrar, srv interface{}) {

}


func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	s.server = grpc.NewServer()

	if err := s.server.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() {
	if s.server != nil {
		s.server.GracefulStop()
	}
}
