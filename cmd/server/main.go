package main

import (
	"flag"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	"fault-tolerant-kv/api/proto"
	"fault-tolerant-kv/internal/grpcserver"
	"fault-tolerant-kv/internal/kv"
	"fault-tolerant-kv/internal/replication"
)

func main() {
	nodeID := flag.Int("id", 1, "node ID")
	port := flag.Int("port", 50051, "gRPC port")

	flag.Parse()

	store := kv.NewStore()

	peers := map[int]string{
                 1: "kv-node1:50051",
                 2: "kv-node2:50052",
                 3: "kv-node3:50053",
	}

	replicator := replication.NewClient(*nodeID, peers)

	server := grpc.NewServer()

	proto.RegisterKVServiceServer(
		server,
		grpcserver.NewServer(store, *nodeID, replicator),
	)

	address := ":" + strconv.Itoa(*port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(
		"Node %d running on %s",
		*nodeID,
		address,
	)

	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
