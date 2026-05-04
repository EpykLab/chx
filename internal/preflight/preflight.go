package preflight

import (
	"os"

	"github.com/spf13/cobra"
)

func ArgsOrError(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		return
	}

	type statter interface {
		Stat() (os.FileInfo, error)
	}

	if in, ok := cmd.InOrStdin().(statter); ok {
		if info, err := in.Stat(); err == nil {
			if info.Mode()&os.ModeCharDevice == 0 {
				return
			}
		}
	}

	_ = cmd.Help()
	os.Exit(1)
}
