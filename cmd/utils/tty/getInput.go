package tty

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"io"

	"github.com/spf13/cobra"
)

func GetInputOrSet(cmd *cobra.Command, args []string) ([]string, error) {

	var input []string

	if len(args) > 0 {
		input = append(input, args[0])
	} else {
		// If no argument is provided, read from stdin
		var inputRead io.Reader = cmd.InOrStdin()
		iter := bufio.NewReader(inputRead)

		for {
			line, err := iter.ReadString('\n')
			if err != nil {
				if err.Error() == "EOF" {
					break
				} else {
					fmt.Fprintln(os.Stderr, "Error reading input:", err)
					return []string{}, err
				}
			}

			l := strings.TrimSpace(string(line))
			input = append(input, l)
		}
	}

	return input, nil
}
