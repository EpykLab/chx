/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/EpykLab/chx/cmd/sources"
	"github.com/spf13/cobra"
)

// vtCmd represents the vt command
var vthashCmd = &cobra.Command{
	Use:   "vthash",
	Short: "Check VirusTotal",
	Long:  `Check a hash against Virus Total`,
	Run: func(cmd *cobra.Command, args []string) {
		var input string
		if len(args) > 0 {
			// If an argument is provided, use it as the input
			input = args[0]
		} else {
			// If no argument is provided, read from stdin
			var inputRead io.Reader = cmd.InOrStdin()
			inputBytes, err := io.ReadAll(inputRead)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error reading input:", err)
				return
			}
			input = string(inputBytes)
		}

		// Trim any extra newlines or spaces from the input
		input = strings.TrimSpace(input)

		// Process the hash
		sources.MakeVtRequest(input)

	},
}

func init() {
	rootCmd.AddCommand(vthashCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
