package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"fault-tolerant-kv/api/proto"
)

func main() {
	nodes := []struct {
		name    string
		address string
	}{
		{"Node 1", "localhost:50051"},
		{"Node 2", "localhost:50052"},
		{"Node 3", "localhost:50053"},
	}

	for _, node := range nodes {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		conn, err := grpc.DialContext(
			ctx,
			node.address,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
		)

		if err != nil {
			fmt.Println(node.name, "UNAVAILABLE:", err)
			cancel()
			continue
		}

		client := proto.NewKVServiceClient(conn)

		if node.name == "Node 1" {
			resp, err := client.Set(ctx, &proto.SetRequest{
				Key:   "name",
				Value: "harshitha",
			})
			if err != nil {
				fmt.Println("Node 1 SET failed:", err)
			} else {
				fmt.Println("Node 1 SET:", resp.Success)
			}
		}

		resp, err := client.Get(ctx, &proto.GetRequest{Key: "name"})
		if err != nil {
			fmt.Println(node.name, "GET failed:", err)
		} else {
			fmt.Println(node.name, "GET:", resp.Value, "FOUND:", resp.Found)
		}

		conn.Close()
		cancel()
	}
}
