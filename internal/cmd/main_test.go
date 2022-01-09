package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"internal/cmd"
)

const help = ""

func TestMain(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		stdin  string
		stdout string
		stderr string
		err    error
	}{
		{
			name:   "NoArgumentsGiven",
			stderr: help,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			cmd.Stdin = strings.NewReader(test.stdin)
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			if have, want := cmd.Main(test.args...), test.err; have != want {
				t.Errorf("err = %q, want %q", have, want)
			}

			if have, want := test.stdout, stdout.String(); have != want {
				t.Errorf("stdout = %q, want %q", have, want)
			}

			if have, want := test.stderr, stderr.String(); have != want {
				t.Errorf("stderr = %q, want %q", have, want)
			}
		})
	}
}
