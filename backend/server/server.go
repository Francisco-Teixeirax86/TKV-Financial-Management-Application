package main

import (
	"backend/data"
	"backend/distributedkv"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var currentTerm int64

type server struct {
	distributedkv.UnimplementedNodeServiceServer
	logger *data.LogStore
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: server <port>")
	}

	port := os.Args[1]
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	logger := data.NewLogStore()

	grpcServer := grpc.NewServer()

	distributedkv.RegisterNodeServiceServer(grpcServer, &server{logger: logger})

	log.Printf("gRPC server listening at %v", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (server *server) SendHeartbeat(ctx context.Context, in *distributedkv.HeartbeatRequest) (*distributedkv.HeartbeatResponse, error) {
	log.Printf("Received heartbeat from leader: %s, term: %d", in.LeaderId, in.Term)
	return &distributedkv.HeartbeatResponse{Success: true, Term: in.Term}, nil
}

func (server *server) RequestVote(ctx context.Context, in *distributedkv.VoteRequest) (*distributedkv.VoteResponse, error) {
	log.Printf("Received vote request from candidate: %s, term: %d", in.CandidateId, in.Term)

	if in.Term >= currentTerm {
		currentTerm = in.Term
		return &distributedkv.VoteResponse{VoteGranted: true, Term: currentTerm}, nil
	}

	return &distributedkv.VoteResponse{VoteGranted: false, Term: currentTerm}, nil
}

func (server *server) AppendEntries(ctx context.Context, in *distributedkv.EntryRequest) (*distributedkv.EntryResponse, error) {
	log.Printf("Received AppendEntries from leader: %s, term: %d", in.LeaderId, in.Term)

	for _, entry := range in.Entries {
		server.logger.AddLogEntry(data.LogEntry{Term: int(entry.Term), Command: entry.Value})
	}
	server.getLogStoreHandler()
	return &distributedkv.EntryResponse{Success: true, Term: in.Term}, nil
}

func (server *server) getLogStoreHandler() {
	for _, logEntry := range server.logger.GetLogs() {
		log.Printf("Log Entry: Term %d, Value: %s", logEntry.Term, logEntry.Command)
	}
}
