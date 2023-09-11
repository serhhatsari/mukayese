package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yildizozan/mukayese/internal"
)

func checkDirsArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
		return err
	}
	if err := cobra.MaximumNArgs(2)(cmd, args); err != nil {
		return err
	}
	return nil
}

func checkDirs(cmd *cobra.Command, args []string) {

	var hashSummaryCurrent map[string]string
	var hashSummaryPrevious map[string]string

	if !internal.IsDirectory(args[0]) {
		log.Fatal("current path is not directory")
	}

	if !internal.IsDirectory(args[1]) {
		log.Fatal("previous path is not directory")
	}

	hashSummaryCurrent = make(map[string]string)
	hashSummaryPrevious = make(map[string]string)

	internal.ListFilesChecksums(hashSummaryCurrent, args[0])
	internal.ListFilesChecksums(hashSummaryPrevious, args[1])

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

// dirsCmd represents the dirs command
var dirsCmd = &cobra.Command{
	Use:   "dirs",
	Short: "Compare two directories",
	Long:  `Compare two directories`,
	Args:  checkDirsArgs,
	Run:   checkDirs,
}

func init() {
	rootCmd.AddCommand(dirsCmd)
}
