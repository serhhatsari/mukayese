package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yildizozan/mukayese/internal"
)

func checkSummaryArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.MinimumNArgs(0)(cmd, args); err != nil {
		return err
	}
	if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
		return err
	}
	return nil
}

func getSummary(cmd *cobra.Command, args []string) {

	var hashSummaryCurrent map[string]string
	hashSummaryCurrent = make(map[string]string)

	var path string
	if 0 == len(args) {
		path = "."
	} else {
		path = args[0]
	}

	internal.ListFilesChecksums(hashSummaryCurrent, path)

	// Open file for output
	output, errFlag := cmd.Flags().GetString("output")
	if errFlag != nil {
		log.Fatalln("parameter oll")
	}

	if output == "" {
		internal.PrintMap("", &hashSummaryCurrent)
	} else {
		success := internal.Export(output, &hashSummaryCurrent)
		if success {
			fmt.Printf("checksums write to %s file\n", output)
		}
	}
}

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Summary",
	Long:  `Summary`,
	Args:  checkSummaryArgs,
	Run:   getSummary,
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	summaryCmd.Flags().StringP("output", "o", "", "output file")
}
