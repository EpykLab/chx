/*
Copyright © 2024 EpykLab contact@epyklab.com
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

// hashCmd represents the hash command
var avhashCmd = &cobra.Command{
	Use:   "avhash",
	Short: "Get details about a hash",
	Long: `Get details about a file hash using Alient Vault. This requires
	an Alien Vault API key`,
	Args: cobra.MaximumNArgs(1), // Allow up to one argument
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
		sources.GetHashDetails(input)
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hashCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hashCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
