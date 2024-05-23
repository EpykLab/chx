/*
Copyright © 2024 EpykLab contact@epyklab.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/EpykLab/ipcheck/cmd/utils"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information about an ip address",
	Long: `Get information from AbuseIP DB about the IP passeded in. IP's can 
	passed in as a argument to get, or, passed to get through stdin.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				utils.GetIPInfo(scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}
		} else {
			utils.GetIPInfo(args[0])
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//getCmd.Flags().String("")
}
