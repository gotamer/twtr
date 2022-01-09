package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"internal/cmd"
)

const help = `Usage: twtxt COMMAND [OPTIONS] [ARGS...]

Decentralized, minimalist microblogging for hackers.

Options:
    -c, --config PATH  Specify a custom config file location.
    -v, --version      Enable verbose output for debugging.
    --version          Show the version and exit.
    -h, --help         Show this message and exit.

Commands:
    timeline   Retrieve your personal timeline.
    following  View the sources that you are following.
    follow     Add a new source to your followings.
    unfollow   Remove an existing source from your list.
    tweet      Send out a message into the void.
    view       View a source that you follow.
    config     Update your configuration.
`

func TestMain(t *testing.T) {
	tests := []struct {
		args   []string
		stdin  string
		stdout string
		stderr string
		err    error
	}{
		{
			stderr: help,
		},
		{
			args:   []string{},
			stderr: help,
		},
		{
			args:   []string{"-h"},
			stderr: help,
		},
		{
			args:   []string{"--help"},
			stderr: help,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(strings.Join(test.args, " "), func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			cmd.Stdin = strings.NewReader(test.stdin)
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			if have, want := cmd.Main(test.args...), test.err; have != want {
				t.Errorf("err = %q, want %q", have, want)
			}

			if have, want := stdout.String(), test.stdout; have != want {
				t.Errorf("stdout = %q, want %q", have, want)
			}

			if have, want := stderr.String(), test.stderr; have != want {
				t.Errorf("stderr = %q, want %q", have, want)
			}
		})
	}
}
