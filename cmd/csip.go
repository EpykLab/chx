/*
Copyright © 2024 contact@epyklab.com
*/

package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/EpykLab/chx/cmd/sources"
	"github.com/spf13/cobra"
)

// researchCmd represents the research command
var csipCmd = &cobra.Command{
	Use:   "csip",
	Short: "Get information about an IP address (cs)",
	Long: `Get information about an IP address using Crowdsec.
	Requires a crowsec API key. Note that Crowdsec levies daily 
	limits on the amount of IP's you can check.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				sources.GetCrowdSecSmoke(scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}
		} else {
			sources.GetCrowdSecSmoke(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(csipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// researchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	csipCmd.Flags().String("hash", "h", "gets information about a hash value. Require OTX api key")
	csipCmd.Flags().String("domain", "d", "gets information about a domain name. Requires OTX api key")
}
