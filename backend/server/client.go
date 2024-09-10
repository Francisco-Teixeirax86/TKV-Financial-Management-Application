package main

import (
	"backend/distributedkv"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	connection, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer connection.Close()

	client := distributedkv.NewNodeServiceClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SendHeartbeat(ctx, &distributedkv.HeartbeatRequest{LeaderId: "leader1", Term: 1})
	if err != nil {
		log.Fatalf("could not send heartbeat: %v", err)
	}

	log.Printf("heartbeat: %v | Success: %v | Term: %v", response, response.GetSuccess(), response.GetTerm())
}
