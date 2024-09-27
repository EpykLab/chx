/*
Copyright © 2024 EpykLab contact@epyklab.com
*/
package cmd

import (
	"github.com/EpykLab/chx/cmd/utils/configs"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure chx",
	Long:  `Configure API connections for chx`,
	Run: func(cmd *cobra.Command, args []string) {
		configs.UpdateConfig()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
