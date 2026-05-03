package preflight

import (
	"os"

	"github.com/spf13/cobra"
)

func ArgsOrError(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(1)
	} else {
		return
	}
}
