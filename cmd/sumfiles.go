/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/yildizozan/mukayese/internal"
)

func checkSum(cmd *cobra.Command, args []string) {
	current := args[0]
	previous := args[1]

	fileCurrent, err := os.Open(current)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer fileCurrent.Close()

	hashSummaryCurrent, err := internal.ParseSumFile(fileCurrent)
	if err != nil {
		log.Fatalf("source sum file err %v", err)
		return
	}

	filePrevious, err := os.Open(previous)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer filePrevious.Close()

	hashSummaryPrevious, err := internal.ParseSumFile(filePrevious)
	if err != nil {
		log.Fatalf("source sum file err %v", err)
		return
	}

	added := make(map[string]string)
	changed := make(map[string]string)
	deleted := make(map[string]string)

	// Determine added and changed files
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

	// Determine deleted files
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

	isEmpty := internal.IsEmpty(added, changed, deleted)
	if isEmpty {
		log.Println("Checksum files are same!")
	} else {
		internal.PrintMap("Added", &added)
		internal.PrintMap("Changed", &changed)
		internal.PrintMap("Deleted", &deleted)
	}

}

var sumfilesCmd = &cobra.Command{
	Use:   "sumfiles",
	Short: "Compare two checksum files",
	Long:  `Compare two checksum files.`,
	Run:   checkSum,
}

func init() {
	rootCmd.AddCommand(sumfilesCmd)
}
