package main

import (
	"fmt"
	"os"

	"fault-tolerant-kv/internal/raft/persist"
)

func main() {
	path := "/tmp/fault-tolerant-raft.log"

	log := persist.New(path)

	if len(os.Args) > 1 && os.Args[1] == "write" {
		err := log.Append(persist.Entry{
			Term:  1,
			Key:   "name",
			Value: "harshitha",
		})

		if err != nil {
			panic(err)
		}

		fmt.Println("Data written to disk")
		return
	}

	entries, err := persist.Recover(path)
	if err != nil {
		panic(err)
	}

	fmt.Println("Recovered entries:")

	for _, entry := range entries {
		fmt.Printf(
			"term=%d key=%s value=%s\n",
			entry.Term,
			entry.Key,
			entry.Value,
		)
	}
}
