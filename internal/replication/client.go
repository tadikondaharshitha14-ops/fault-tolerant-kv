package replication

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"fault-tolerant-kv/api/proto"
)

type Client struct {
	nodeID int
	peers  map[int]string
}

func NewClient(nodeID int, peers map[int]string) *Client {
	return &Client{
		nodeID: nodeID,
		peers:  peers,
	}
}

func (c *Client) ReplicateSet(key, value string) int {
	success := 0

	for id, address := range c.peers {
		if id == c.nodeID {
			continue
		}

		if c.replicate(id, address, key, value, "SET") {
			success++
		}
	}

	return success
}

func (c *Client) ReplicateDelete(key string) int {
	success := 0

	for id, address := range c.peers {
		if id == c.nodeID {
			continue
		}

		if c.replicate(id, address, key, "", "DELETE") {
			success++
		}
	}

	return success
}

func (c *Client) replicate(
	nodeID int,
	address string,
	key string,
	value string,
	operation string,
) bool {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		fmt.Printf(
			"Node %d replication failed: %v\n",
			nodeID,
			err,
		)
		return false
	}
	defer conn.Close()

	client := proto.NewKVServiceClient(conn)

	response, err := client.Replicate(
		ctx,
		&proto.ReplicateRequest{
			Key:       key,
			Value:     value,
			Operation: operation,
		},
	)

	if err != nil {
		fmt.Printf(
			"Node %d replication failed: %v\n",
			nodeID,
			err,
		)
		return false
	}

	return response.GetSuccess()
}
