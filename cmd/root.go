/*
Copyright © 2024 EpykLab contact@epyklab.com
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string
var version string

var rootCmd = &cobra.Command{
	Use:     "chx",
	Version: version,
	Short:   "IP addr research in the command line",
	Long: `chx is a cli tool that aims to streamline workflows by bringing IP 
address lookup and research to the command line, making use of stdin and stdout, 
allowing it to integrate seemlessly with other tools`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
