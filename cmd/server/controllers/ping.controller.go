package controllers

import (
	"context"

	"github.com/MrAzharuddin/grpc_tutorial/pb"
)

type PingServer struct {
	pb.UnimplementedPingServiceServer
}

func NewPingServer() *PingServer{
	return &PingServer{}
}

func (ps *PingServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Message: "Hello! It's me :) Your server"}, nil
}  
	
