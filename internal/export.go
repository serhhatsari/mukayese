package internal

import (
	"fmt"
	"log"
	"os"
)

func Export(filename string, checksums *map[string]string) bool {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("sum file not created\n %v", err)
		return false
	}
	defer f.Close()

	for key, val := range *checksums {
		fmt.Fprintf(f, "%s@sha256:%s\n", key, val)
	}

	return true
}
