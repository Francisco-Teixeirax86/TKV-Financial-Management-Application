package main

import (
	"backend/distributedkv"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	distributedkv.UnimplementedNodeServiceServer
}

func (server *server) SendHeartbeat(ctx context.Context, in *distributedkv.HeartbeatRequest) (*distributedkv.HeartbeatResponse, error) {
	log.Printf("Received heartbeat from leader: %s, term: %d", in.LeaderId, in.Term)
	return &distributedkv.HeartbeatResponse{Success: true, Term: in.Term}, nil
}

func main() {
	listner, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	distributedkv.RegisterNodeServiceServer(grpcServer, &server{})

	log.Printf("gRPC server listening at %v", listner.Addr())
	if err := grpcServer.Serve(listner); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
