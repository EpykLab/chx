/*
Copyright © 2024 contact@epyklab.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/EpykLab/chx/cmd/sources"
	"github.com/EpykLab/chx/cmd/utils/pretty"
	"github.com/EpykLab/chx/cmd/utils/pretty/data"
	"github.com/EpykLab/chx/cmd/utils/shared"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// domainCmd represents the domain command
var avdomainCmd = &cobra.Command{
	Use:   "alienvault",
	Short: "Get details about a domain name",
	Long: `Get details about a domain name from Alient Vault. Requirs
	an AlienVault API Key`,
	Run: func(cmd *cobra.Command, args []string) {
		formated := cmd.Flag("format").Changed

		var result interface{}

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
		result = sources.GetDomainDetailsAV(input)

		if formated {
			err := pretty.PrintContentPretty(data.Domain, data.AlienVault, result)
			if err != nil {
				log.Error(err)
			}
		} else {
			shared.Out(result)
		}

	},
}

func init() {
	avdomainCmd.Flags().Bool("format", false, "pretty print results")
}
