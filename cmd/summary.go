/*
Copyright © 2023 Ozan YILDIZ <developer@yildizozan.com>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yildizozan/mukayese/internal"
	"log"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(0)(cmd, args); err != nil {
			return err
		}
		if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

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
			for key, val := range hashSummaryCurrent {
				fmt.Printf("%s@sha256:%s\n", key, val)
			}
		} else {
			success := internal.Export(output, &hashSummaryCurrent)
			if success {
				fmt.Printf("checksums write to %s file\n", output)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// summaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// summaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	summaryCmd.Flags().StringP("output", "o", "", "output file")
}
