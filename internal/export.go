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

	sortedDatas := Sort(checksums)

	for i := 0; i < len(*sortedDatas); i++ {
		fmt.Fprintf(f, "%s\n", (*sortedDatas)[i])
	}

	return true
}
