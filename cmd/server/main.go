package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"fault-tolerant-kv/api/proto"
	"fault-tolerant-kv/internal/grpcserver"
	"fault-tolerant-kv/internal/kv"
)

func main() {
	store := kv.NewStore()

	server := grpc.NewServer()

	proto.RegisterKVServiceServer(
		server,
		grpcserver.NewServer(store),
	)

	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("KV gRPC server running on :50051")

	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
