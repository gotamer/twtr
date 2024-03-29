package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const helpMessage = `Usage: twtr COMMAND [OPTIONS] [ARGS...]

Decentralized, minimalist microblogging for hackers.

Options:
	-c, --config PATH  Specify a custom config file location.
	-h, --help         Show this message and exit.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.

Commands:
	quickstart  Quickstart wizard for setting up twtxt.
	timeline    Retrieve your personal timeline.
	following   View the sources that you are following.
	follow      Add a new source to your followings.
	unfollow    Remove an existing source from your list.
	tweet       Send out a message into the void.
	view        View a source that you follow.
	config      Update your configuration.
`

func TestMain(t *testing.T) {
	tests := []struct {
		args   []string
		stdin  string
		stdout string
		stderr string
		err    error
	}{
		// no subcommand
		{
			stderr: "twtr: no COMMAND given\n\n" + helpMessage,
		},
		{
			args:   []string{},
			stderr: "twtr: no COMMAND given\n\n" + helpMessage,
		},
		{
			args:   []string{"-h"},
			stderr: helpMessage,
		},
		{
			args:   []string{"--help"},
			stderr: helpMessage,
		},

		// quickstart
		{
			args:   []string{"quickstart", "-h"},
			stderr: quickstartCommand.help(&Context{Self: "twtr"}),
		},
		{
			args:   []string{"quickstart", "--help"},
			stderr: quickstartCommand.help(&Context{Self: "twtr"}),
		},

		// timeline
		{
			args:   []string{"timeline", "-h"},
			stderr: timelineCommand.help(&Context{Self: "twtr"}),
		},
		{
			args:   []string{"timeline", "--help"},
			stderr: timelineCommand.help(&Context{Self: "twtr"}),
		},

		// following
		{
			args:   []string{"following", "-h"},
			stderr: followingCommand.help(&Context{Self: "twtr"}),
		},
		{
			args:   []string{"following", "--help"},
			stderr: followingCommand.help(&Context{Self: "twtr"}),
		},

		// follow
		{
			args:   []string{"follow"},
			stderr: "twtr follow: no SOURCE given\n\n" + followCommand.help(&Context{Self: "twtr"}),
		},
		{
			args:   []string{"follow", "-h"},
			stderr: followCommand.help(&Context{Self: "twtr"}),
		},
		{
			args:   []string{"follow", "--help"},
			stderr: followCommand.help(&Context{Self: "twtr"}),
		},

		// unfollow
		{
			args:   []string{"unfollow"},
			stderr: "twtr unfollow: no SOURCE given\n\n" + unfollowCommand.help(&Context{Self: "twtr"}),
		},
		{
			args:   []string{"unfollow", "-h"},
			stderr: unfollowCommand.help(&Context{Self: "twtr"}),
		},
		{
			args:   []string{"unfollow", "--help"},
			stderr: unfollowCommand.help(&Context{Self: "twtr"}),
		},

		// tweet
		{
			args:   []string{"tweet"},
			stderr: "twtr tweet: no TWEET given\n\n" + tweetCommand.help(&Context{Self: "twtr"}),
		},
		{
			args:   []string{"tweet", "-h"},
			stderr: tweetCommand.help(&Context{Self: "twtr"}),
		},
		{
			args:   []string{"tweet", "--help"},
			stderr: tweetCommand.help(&Context{Self: "twtr"}),
		},

		// view
		{
			args:   []string{"view"},
			stderr: "twtr view: no SOURCE given\n\n" + viewCommand.help(&Context{Self: "twtr"}),
		},
		{
			args:   []string{"view", "-h"},
			stderr: viewCommand.help(&Context{Self: "twtr"}),
		},
		{
			args:   []string{"view", "--help"},
			stderr: viewCommand.help(&Context{Self: "twtr"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(strings.Join(test.args, " "), func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			ctx := Context{
				Stdin:  strings.NewReader(test.stdin),
				Stdout: &stdout,
				Stderr: &stderr,
			}

			if have, want := Main(&ctx, test.args...), test.err; have != want {
				t.Errorf("err = %q, want %q", have, want)
			}

			if have, want := stdout.String(), test.stdout; have != want {
				t.Errorf("stdout diff:\n%s", cmp.Diff(have, want))
			}

			if have, want := stderr.String(), test.stderr; have != want {
				t.Errorf("stderr diff:\n%s", cmp.Diff(have, want))
			}
		})
	}
}
