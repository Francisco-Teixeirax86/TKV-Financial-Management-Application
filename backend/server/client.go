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
	nodes := []string{"localhost:50051", "localhost:50052", "localhost:50053"}

	for _, node := range nodes {
		conn, client, err := connectToNode(node)
		if err != nil {
			log.Fatalf("failed to connect to node %s: %v", node, err)
		}
		defer conn.Close()

		// Simulate leader sending heartbeat and requesting votes
		sendHeartbeat(client, node)
		requestVote(client, node)
		appendEntries(client, node)
	}
}

func connectToNode(nodeAddress string) (*grpc.ClientConn, distributedkv.NodeServiceClient, error) {
	conn, err := grpc.NewClient("passthrough:"+nodeAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	client := distributedkv.NewNodeServiceClient(conn)
	return conn, client, nil
}

func requestVote(c distributedkv.NodeServiceClient, nodeAddress string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Printf("Requesting vote from node: %s", nodeAddress)
	r, err := c.RequestVote(ctx, &distributedkv.VoteRequest{CandidateId: "node1", Term: 1})
	if err != nil {
		log.Fatalf("could not request vote: %v", err)
	}
	log.Printf("Vote granted: %v from node: %s", r.VoteGranted, nodeAddress)
}

func sendHeartbeat(c distributedkv.NodeServiceClient, nodeAddress string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Printf("Sending heartbeat to node: %s", nodeAddress)
	r, err := c.SendHeartbeat(ctx, &distributedkv.HeartbeatRequest{LeaderId: "leader1", Term: 1})
	if err != nil {
		log.Fatalf("could not send heartbeat: %v", err)
	}
	log.Printf("Heartbeat success: %v from node: %s", r.Success, nodeAddress)
}

func appendEntries(client distributedkv.NodeServiceClient, nodeAddr string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	logEntries := []*distributedkv.LogEntry{
		{Value: "transaction1", Term: 1},
		{Value: "transaction2", Term: 1},
	}

	response, err := client.AppendEntries(ctx, &distributedkv.EntryRequest{
		LeaderId:     "leader1",
		Term:         1,
		Entries:      logEntries,
		LeaderCommit: 1,
	})
	if err != nil {
		log.Fatalf("could not send append entries: %v", err)
	}

	log.Printf("AppendEntries response success: %v from node: %s", response.GetSuccess(), nodeAddr)
}
