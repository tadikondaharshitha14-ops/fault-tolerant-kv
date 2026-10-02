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

func main() {
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := proto.NewKVServiceClient(conn)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	setResp, err := client.Set(ctx, &proto.SetRequest{
		Key:   "name",
		Value: "harshitha",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SET:", setResp.Success)

	getResp, err := client.Get(ctx, &proto.GetRequest{
		Key: "name",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GET:", getResp.Value)
	fmt.Println("FOUND:", getResp.Found)

	deleteResp, err := client.Delete(ctx, &proto.DeleteRequest{
		Key: "name",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DELETE:", deleteResp.Success)
}
