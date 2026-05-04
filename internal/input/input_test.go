package input

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestGetInputOrSet_UsesArgWhenProvided(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.SetIn(strings.NewReader("stdin-value\n"))

	got, err := GetInputOrSet(cmd, []string{"from-arg"})
	if err != nil {
		t.Fatalf("GetInputOrSet() error = %v", err)
	}
	if len(got) != 1 || got[0] != "from-arg" {
		t.Fatalf("unexpected result: %#v", got)
	}
}

func TestGetInputOrSet_ReadsStdinLines(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.SetIn(strings.NewReader("one\n two \n\nthree\n"))

	got, err := GetInputOrSet(cmd, nil)
	if err != nil {
		t.Fatalf("GetInputOrSet() error = %v", err)
	}

	want := []string{"one", "two", "", "three"}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d (%#v)", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("line[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func FuzzGetInputOrSet(f *testing.F) {
	f.Add("a\nb\n")
	f.Add("single-line")
	f.Add("\n\n")

	f.Fuzz(func(t *testing.T, in string) {
		cmd := &cobra.Command{}
		cmd.SetIn(strings.NewReader(in))

		got, err := GetInputOrSet(cmd, nil)
		if err != nil {
			t.Fatalf("GetInputOrSet() error = %v", err)
		}

		for i, line := range got {
			if strings.Contains(line, "\n") {
				t.Fatalf("line[%d] still contains newline: %q", i, line)
			}
		}
	})
}
