/*
Copyright © 2024 EpykLab contact@epyklab.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/EpykLab/chx/cmd/sources"
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
		if len(args) < 1 {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				sources.GetIPDetails(scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}
		} else {
			sources.GetIPDetails(args[0])
		}

	},
}

func init() {}
