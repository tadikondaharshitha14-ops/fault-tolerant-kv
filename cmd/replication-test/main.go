package main

import (
	"fmt"

	"fault-tolerant-kv/internal/replication"
)

func main() {
	peers := map[int]string{
		1: "localhost:50051",
		2: "localhost:50052",
		3: "localhost:50053",
	}

	r := replication.NewClient(1, peers)

	fmt.Println("Successful replicas:", r.ReplicateSet("replica-test", "hello"))
}
