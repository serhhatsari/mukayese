/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var hashSummaryCurrent map[string]string

var hashSummaryPrevious map[string]string

// isDirectory determines if a file represented
// by `path` is a directory or not
func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

func Checksum(path string) string {
	hasher := sha256.New()
	s, err := os.ReadFile(path)
	hasher.Write(s)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

func listFilesChecksums(summary map[string]string, path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	//infos := make([]fs.FileInfo, 0, len(entries))

	for _, entry := range entries {
		//fmt.Println(entry.Name())

		newPath := filepath.Join(path, entry.Name())

		// project/dll/hello.cs -> dll/hello.cs
		x := strings.Split(newPath, string(os.PathSeparator))
		var y string
		for _, val := range x[1:] {
			y = filepath.Join(y, val)
		}

		// check is dir
		if !entry.IsDir() {
			summary[y] = Checksum(newPath)
			//fmt.Printf("%s -> %s\n", newPath, checksum)

		} else {
			listFilesChecksums(summary, newPath)
		}

	}
}

func main() {
	//cmd.Execute()

	// Usage
	if len(os.Args) == 2 {
		fmt.Println("mukayese [CURRENT] [PREVIOUS]")
		os.Exit(1)
	}

	if !isDirectory(os.Args[1]) {
		log.Fatal("current path is not directory")
	}

	if !isDirectory(os.Args[2]) {
		log.Fatal("previous path is not directory")
	}

	hashSummaryCurrent = make(map[string]string)
	hashSummaryPrevious = make(map[string]string)

	listFilesChecksums(hashSummaryCurrent, os.Args[1])
	listFilesChecksums(hashSummaryPrevious, os.Args[2])

	fmt.Printf("Current: \n")
	for key, val := range hashSummaryCurrent {
		fmt.Printf("%s@sha256:%s\n", key, val)
	}

	fmt.Printf("Previous: \n")
	for key, val := range hashSummaryPrevious {
		fmt.Printf("%s@sha256:%s\n", key, val)
	}

	added := make(map[string]string)
	changed := make(map[string]string)
	deleted := make(map[string]string)

	for currKey, currVal := range hashSummaryCurrent {
		exist := true
		for prevKey, prevVal := range hashSummaryPrevious {
			if currKey == prevKey {
				exist = false
				if currVal != prevVal {
					changed[currKey] = currVal
				}
			}
		}
		if exist {
			added[currKey] = currVal
		}
	}

	// Determine deteled files
	for prevKey, prevVal := range hashSummaryPrevious {
		exist := true
		for currKey := range hashSummaryCurrent {
			if prevKey == currKey {
				exist = false
				break
			}
		}
		if exist {
			deleted[prevKey] = prevVal
		}
	}

	fmt.Println() // Temp

	// Added
	fmt.Printf("Added: \n")
	for key, val := range added {
		fmt.Printf("%s@sha256:%s\n", key, val)
	}
	fmt.Println()

	// Changed
	fmt.Printf("Changed: \n")
	for key, val := range changed {
		fmt.Printf("%s@sha256:%s\n", key, val)
	}
	fmt.Println()

	// Deleted
	fmt.Printf("Deleted: \n")
	for key, val := range deleted {
		fmt.Printf("%s@sha256:%s\n", key, val)
	}

}
