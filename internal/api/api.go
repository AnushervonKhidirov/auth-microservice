package api

import (
	"net"

	pb "auth/gen/go"
	"auth/internal/config"

	"google.golang.org/grpc"
)

type Server struct {
	Network string
	Address string
}

func NewServer(conf config.GRPCServer) *Server {
	return &Server{Network: conf.Network, Address: conf.Address}
}

func (s Server) Serve() error {
	listener, err := net.Listen(s.Network, s.Address)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &AuthServiceServer{})
	pb.RegisterUserServiceServer(grpcServer, &UserServiceServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}
