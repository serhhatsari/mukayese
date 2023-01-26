/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yildizozan/mukayese/internal"
	"log"
	"os"
)

// sumfilesCmd represents the sumfiles command
var sumfilesCmd = &cobra.Command{
	Use:   "sumfiles",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
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

		isEmpty := internal.IsEmpty(added, changed, deleted)
		if isEmpty {
			log.Println("Checksum files are same!")
		} else {
			// Deleted
			internal.PrintMap("Added", &added)
			internal.PrintMap("Changed", &changed)
			internal.PrintMap("Deleted", &deleted)
		}

	},
}

func init() {
	rootCmd.AddCommand(sumfilesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sumfilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sumfilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
