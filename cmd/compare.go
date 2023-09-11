package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/yildizozan/mukayese/internal"
)

// check arguments for compare command
// 2 arguments are required
func checkArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
		return err
	}
	if err := cobra.MaximumNArgs(2)(cmd, args); err != nil {
		return err
	}

	if !internal.IsDirectory(args[0]) {
		return fmt.Errorf("current path is not directory")
	}

	if !internal.IsDirectory(args[1]) {
		return fmt.Errorf("previous path is not directory")
	}

	return nil
}

func compareFolders(cmd *cobra.Command, args []string) {

	hashSummaryCurrent := make(map[string]string)
	hashSummaryPrevious := make(map[string]string)

	internal.ListFilesChecksums(hashSummaryCurrent, args[0])
	internal.ListFilesChecksums(hashSummaryPrevious, args[1])

	color.Cyan("-- Comparing %s and %s --\n\n", args[0], args[1])

	color.Cyan("Current: \n")
	for key, val := range hashSummaryCurrent {
		fmt.Printf("%s@sha256:%s\n", key, val)
	}

	color.Cyan("\nPrevious: \n")
	for key, val := range hashSummaryPrevious {
		fmt.Printf("%s@sha256:%s\n", key, val)
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

	color.Cyan("\n\n-- RESULTS --\n")

	color.Green("\nAdded: \n")
	for key, val := range added {
		fmt.Printf("%s@sha256:%s\n", key, val)
	}
	fmt.Println()

	color.Yellow("Changed: \n")
	for key, val := range changed {
		fmt.Printf("%s@sha256:%s\n", key, val)
	}
	fmt.Println()

	color.Red("Deleted: \n")
	for key, val := range deleted {
		fmt.Printf("%s@sha256:%s\n", key, val)
	}
}

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compare directories.",
	Long:  `Compare two directory to find added, changed and deleted files.`,
	Args:  checkArgs,
	Run:   compareFolders,
}

func init() {
	rootCmd.AddCommand(compareCmd)
	compareCmd.Flags().BoolP("files", "f", false, "Compare files")
}
