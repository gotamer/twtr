package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"internal/cmd"
)

const (
	sourceUsage = `Sources:
	At least one SOURCE must be given (unless called with -h), each SOURCE
	consists of a NICK and a URL. Allowed formats are NICK@URL or NICK URL, if
	you don't know the nickname of a SOURCE, you can make one up, or use the
	domain part of the URL (this can be easily changed later).
`

	help = `Usage: twtr COMMAND [OPTIONS] [ARGS...]

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

	quickstartUsage = `Usage: twtr quickstart [-cfhnuv] [--disclose-identity] [--follow-news]

Quickstart wizard for setting up twtxt.

Options:
	-c, --config PATH        Specify a custom configuration file location.
	    --disclose-identity  Show your nickname and url in the User Agent.
	-f, --file PATH          Specify a custom twtxt file location.
	    --follow-news        Follow the official twtxt and twtr news feeds.
	-h, --help               Show this message and exit.
	-n, --nick NICK          Specify the nickname for your feed.
	-u, --url URL            Specify the url that your feed will be hosted at.
	-v, --verbose            Enable verbose output for debugging.
	    --version            Show the version and exit.
`

	timelineUsage = `Usage: twtr timeline [-chv] [--limit COUNT] [--sort ascending | descending]

Retrieve your personal timeline.

Options:
	-c, --config PATH     Specify a custom configuration file location.
	-h, --help            Show this message and exit.
	    --limit COUNT     Limit the amount of tweets shown.
	    --sort DIRECTION  Sort tweets ascending or descending by timestamp.
	-v, --verbose         Enable verbose output for debugging.
	    --version         Show the version and exit.
`

	followingUsage = `Usage: twtr following [-chv]

View the sources that you are following.

Options:
	-c, --config PATH  Specify a custom configuration file location.
	-h, --help         Show this message and exit.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.
`

	followUsage = `Usage: twtr follow [-chv] [--replace] SOURCE [SOURCES...]

Add a new source to your following.

Options:
	-c, --config PATH  Specify a custom configuration file location.
	-h, --help         Show this message and exit.
	    --replace      Replace duplicates instead of returning error.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.

` + sourceUsage

	unfollowUsage = `Usage: twtr unfollow [-chv] NICK|URL

Remove an existing source from your list.

Options:
	-c, --config PATH  Specify a custom configuration file location.
	-h, --help         Show this message and exit.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.
`

	tweetUsage = `Usage: twtr tweet [-cfhv] TWEET

Send out a message into the void.

Options:
	-c, --config PATH  Specify a custom configuration file location.
	-f, --file PATH    Specify a custom twtxt file location.
	-h, --help         Show this message and exit.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.
`

	viewUsage = `Usage: twtr view [-chv] SOURCE [SOURCES]

View a source that you follow.

Options:
	-c, --config PATH  Specify a custom configuration file location.
	-h, --help         Show this message and exit.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.

` + sourceUsage

	configUsage = `Usage: twtr config [-chv] [--edit] [--remove KEY]|[KEY [VALUE]]

Update your configuration.

Options:
	-c, --config PATH  Specify a custom configuration file location.
	    --edit         Edit the configuration file manually.
	-h, --help         Show this message and exit.
	    --remove KEY   Remove a configuration by its KEY, e.g. twtxt.nick.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.
`
)

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
		{
			args:   []string{"quickstart", "-h"},
			stderr: quickstartUsage,
		},
		{
			args:   []string{"quickstart", "--help"},
			stderr: followingUsage,
		},
		{
			args:   []string{"timeline", "-h"},
			stderr: followingUsage,
		},
		{
			args:   []string{"timeline", "--help"},
			stderr: followingUsage,
		},
		{
			args:   []string{"following", "-h"},
			stderr: followingUsage,
		},
		{
			args:   []string{"following", "--help"},
			stderr: followingUsage,
		},
		{
			args:   []string{"follow"},
			stderr: followUsage,
		},
		{
			args:   []string{"follow", "-h"},
			stderr: followUsage,
		},
		{
			args:   []string{"follow", "--help"},
			stderr: followUsage,
		},
		{
			args:   []string{"unfollow"},
			stderr: unfollowUsage,
		},
		{
			args:   []string{"unfollow", "-h"},
			stderr: unfollowUsage,
		},
		{
			args:   []string{"unfollow", "--help"},
			stderr: unfollowUsage,
		},
		{
			args:   []string{"tweet"},
			stderr: tweetUsage,
		},
		{
			args:   []string{"tweet", "-h"},
			stderr: tweetUsage,
		},
		{
			args:   []string{"tweet", "--help"},
			stderr: tweetUsage,
		},
		{
			args:   []string{"view"},
			stderr: viewUsage,
		},
		{
			args:   []string{"view", "-h"},
			stderr: viewUsage,
		},
		{
			args:   []string{"view", "--help"},
			stderr: viewUsage,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(strings.Join(test.args, " "), func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			ctx := cmd.Context{
				Self:   "twtr",
				Stdin:  strings.NewReader(test.stdin),
				Stdout: &stdout,
				Stderr: &stderr,
			}

			if have, want := cmd.Main(&ctx, test.args...), test.err; have != want {
				t.Errorf("err = %q, want %q", have, want)
			}

			if have, want := stdout.String(), test.stdout; have != want {
				t.Errorf("\nstdout:\n%s\n\nwant:\n%s\n", have, want)
			}

			if have, want := stderr.String(), test.stderr; have != want {
				t.Errorf("\nstderr:\n%s\n\nwant:\n%s\n", have, want)
			}
		})
	}
}
