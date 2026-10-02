package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"fault-tolerant-kv/api/proto"
)

func connect(address string) (*grpc.ClientConn, proto.KVServiceClient) {
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	return conn, proto.NewKVServiceClient(conn)
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	// Node 1
	conn1, node1 := connect("localhost:50051")
	defer conn1.Close()

	setResp, err := node1.Set(ctx, &proto.SetRequest{
		Key:   "name",
		Value: "harshitha",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Node 1 SET:", setResp.Success)

	get1, err := node1.Get(ctx, &proto.GetRequest{
		Key: "name",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Node 1 GET:", get1.Value, "FOUND:", get1.Found)

	// Node 2
	conn2, node2 := connect("localhost:50052")
	defer conn2.Close()

	get2, err := node2.Get(ctx, &proto.GetRequest{
		Key: "name",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Node 2 GET:", get2.Value, "FOUND:", get2.Found)

	// Node 3
	conn3, node3 := connect("localhost:50053")
	defer conn3.Close()

	get3, err := node3.Get(ctx, &proto.GetRequest{
		Key: "name",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Node 3 GET:", get3.Value, "FOUND:", get3.Found)
}
