/*
Copyright © 2024 contact@epyklab.com
*/

package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/EpykLab/ipcheck/cmd/sources"
	"github.com/spf13/cobra"
)

// researchCmd represents the research command
var researchCmd = &cobra.Command{
	Use:   "research",
	Short: "Aggregates responses from sources for a given IP",
	Long: `Research aggregates information from all configured sources into a condensed report

Research can read in from Stdin, or, can read from arguments passed directly. Additionaly, research
can output to several formats include json, or txt, or markdown. If no switches are passed in it defaults
	search ip addresses`,
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
	rootCmd.AddCommand(researchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// researchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	researchCmd.Flags().String("hash", "h", "gets information about a hash value. Require OTX api key")
	researchCmd.Flags().String("domain", "d", "gets information about a domain name. Requires OTX api key")
}
