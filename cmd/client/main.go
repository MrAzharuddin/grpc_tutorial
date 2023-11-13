package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MrAzharuddin/grpc_tutorial/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddress = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	// Create a PingService client
	client := pb.NewPingServiceClient(conn)

	// Call Ping RPC
	pingRequest := &pb.PingRequest{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	pingResponse, err := client.Ping(ctx, pingRequest)
	if err != nil {
		log.Fatalf("Error calling Ping RPC: %v", err)
	}

	// Print the response
	fmt.Println("Server Response:", pingResponse.Message)

}
