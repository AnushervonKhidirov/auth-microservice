package api

import (
	pb "auth/gen/go"
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s AuthServiceServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.JWTResponse, error) {
	fmt.Println(req)
	return nil, nil
}

func (s AuthServiceServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.JWTResponse, error) {
	return nil, nil
}

func (s AuthServiceServer) SignOut(ctx context.Context, req *pb.SignOutRequest) (*emptypb.Empty, error) {
	return nil, nil
}
