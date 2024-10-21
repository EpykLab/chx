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
	"bufio"
	"fmt"
	"os"

	"github.com/EpykLab/chx/cmd/sources"
	"github.com/EpykLab/chx/cmd/utils/pretty"
	"github.com/EpykLab/chx/cmd/utils/pretty/data"
	"github.com/EpykLab/chx/cmd/utils/shared"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var avipCmd = &cobra.Command{
	Use:   "alienvault",
	Short: "Get information about an ip address using Alien Vault",
	Long: `Get information about an IP address using Alient Vault. 
	Requires Alient Vault API key`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		formated := cmd.Flag("format").Changed

		var result interface{}

		if len(args) < 1 {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				result = sources.GetIPDetails(scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}
		} else {
			result = sources.GetIPDetails(args[0])
		}

		if formated {
			err := pretty.PrintContentPretty(data.IP, data.AlienVault, result)
			if err != nil {
				log.Error(err)
			}
		} else {
			shared.Out(result)
		}
	},
}

func init() {
	avipCmd.Flags().Bool("format", false, "pretty print results")
}
