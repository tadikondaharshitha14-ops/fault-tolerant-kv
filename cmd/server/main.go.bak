package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"fault-tolerant-kv/internal/kv"
)

func main() {
	store := kv.NewStore()

	fmt.Println("=================================")
	fmt.Println(" Fault-Tolerant KV Store")
	fmt.Println("=================================")
	fmt.Println("Commands:")
	fmt.Println("  SET key value")
	fmt.Println("  GET key")
	fmt.Println("  DELETE key")
	fmt.Println("  EXIT")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		parts := strings.SplitN(line, " ", 3)

		switch strings.ToUpper(parts[0]) {

		case "SET":
			if len(parts) != 3 {
				fmt.Println("Usage: SET key value")
				continue
			}

			store.Set(parts[1], parts[2])
			fmt.Println("OK")

		case "GET":
			if len(parts) != 2 {
				fmt.Println("Usage: GET key")
				continue
			}

			value, ok := store.Get(parts[1])

			if !ok {
				fmt.Println("NOT_FOUND")
			} else {
				fmt.Println(value)
			}

		case "DELETE":
			if len(parts) != 2 {
				fmt.Println("Usage: DELETE key")
				continue
			}

			store.Delete(parts[1])
			fmt.Println("OK")

		case "EXIT":
			fmt.Println("Bye!")
			return

		default:
			fmt.Println("Unknown command")
		}
	}
}
