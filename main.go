package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/MrAzharuddin/grpc_tutorial/cmd/server/controllers"
	"github.com/MrAzharuddin/grpc_tutorial/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	host = "localhost"
	port = ":50051"
)

func main() {
	// merging host and port
	addr := fmt.Sprintf("%s%s", host, port)

	// create a network listener and handle errors
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("error while starting tcp listener: ", err)
		os.Exit(1)
	}
	log.Println("server started at: ", "http://"+addr)

	// create a new gRPC server
	grpcServer := grpc.NewServer()

	// get controllers of services
	pingCtl := controllers.NewPingServer()

	// 
	pb.RegisterPingServiceServer(grpcServer, pingCtl)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error serving gRPC: ", err)
		os.Exit(1)
	}

}
