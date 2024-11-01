package tty

import (
	"fmt"
	"os"
	"strings"

	"io"

	"github.com/spf13/cobra"
)

func GetInputOrSet(cmd *cobra.Command, args []string) (string, error) {

	var input string

	if len(args) > 0 {
		input = args[0]
	} else {

		// If no argument is provided, read from stdin
		var inputRead io.Reader = cmd.InOrStdin()
		inputBytes, err := io.ReadAll(inputRead)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return "", err
		}
		input = strings.TrimSpace(string(inputBytes))
	}

	return input, nil
}
