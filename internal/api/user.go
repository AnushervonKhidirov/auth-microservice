package api

import (
	pb "auth/gen/go"
	"context"
	"fmt"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	fmt.Println(req)
	return nil, nil
}
