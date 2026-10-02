package storage

import (
	"bufio"
	"os"
	"strings"
)

func readSSTable(file *os.File) (map[string]string, error) {
	result := make(map[string]string)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "\t", 2)

		if len(parts) != 2 {
			continue
		}

		result[parts[0]] = parts[1]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
