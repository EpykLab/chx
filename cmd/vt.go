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

// vtCmd represents the vt command
var vthashCmd = &cobra.Command{
	Use:   "vthash",
	Short: "Check VirusTotal",
	Long:  `Check a hash against Virus Total`,
	Run: func(cmd *cobra.Command, args []string) {
		var input string
		if len(args) > 0 {
			input = args[0]
		} else {
			var inputRead io.Reader = cmd.InOrStdin()
			inputBytes, err := io.ReadAll(inputRead)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error reading input:", err)
				return
			}
			input = string(inputBytes)
		}

		input = strings.TrimSpace(input)

		sources.MakeVtRequest(input)

	},
}

func init() {
}
