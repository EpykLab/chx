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
	"os"

	"github.com/EpykLab/chx/cmd/sources"
	"github.com/EpykLab/chx/cmd/utils/pretty"
	"github.com/EpykLab/chx/cmd/utils/pretty/data"
	"github.com/EpykLab/chx/cmd/utils/shared"
	"github.com/EpykLab/chx/cmd/utils/tty"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var aipdbCmd = &cobra.Command{
	Use:   "aipdb",
	Short: "Search IP address against Abuse IP DB",
	Long: `Use Abuse IP db to search for information about
addresses`,
	Run: func(cmd *cobra.Command, args []string) {
		formated := cmd.Flag("format").Changed

		input, err := tty.GetInputOrSet(cmd, args)
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		result := sources.GetIPInfoIpabd(input)

		if formated {
			err := pretty.PrintContentPretty(data.IP, data.IpAbuseDB, result)
			if err != nil {
				log.Error(err)
			}
		} else {
			shared.Out(result)
		}
	},
}

func init() {
	aipdbCmd.Flags().Bool("format", false, "pretty print results")
}
