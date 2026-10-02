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
)

func main() {
	nodeID := flag.Int("id", 1, "node ID")
	port := flag.Int("port", 50051, "gRPC port")

	flag.Parse()

	store := kv.NewStore()

	server := grpc.NewServer()

	proto.RegisterKVServiceServer(
		server,
		grpcserver.NewServer(store, *nodeID),
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
