package cmd

import (
	"errors"
	"fmt"
	"io"

	"internal/cmd/config"
	"internal/cmd/follow"
	"internal/cmd/following"
	"internal/cmd/quickstart"
	"internal/cmd/timeline"
	"internal/cmd/tweet"
	"internal/cmd/unfollow"
	"internal/cmd/view"
)

const version = "v0.0.0"

var (
	conf    string = "~/.config/twtxt/config"
	verbose bool
)

type command struct {
	main func(args ...string) error
	help func(self, desc string)
	desc string
}

var commands map[string]command = map[string]command{
	"config": {
		main: config.Main,
		desc: "Update your configuration.",
	},
	"follow": {
		main: follow.Main,
		desc: "Add a new source to your followings.",
	},
	"following": {
		main: following.Main,
		desc: "View the sources that you are following.",
	},
	"quickstart": {
		main: quickstart.Main,
		desc: "Setup a new config file with the quickstart wizard.",
	},
	"timeline": {
		main: timeline.Main,
		desc: "Retrieve your personal timeline.",
	},
	"tweet": {
		main: tweet.Main,
		desc: "Send out a message into the void.",
	},
	"unfollow": {
		main: unfollow.Main,
		desc: "Remove an existing source from your list.",
	},
	"view": {
		main: view.Main,
		desc: "View a source that you follow.",
	},
}

func help(self string) string {
	usage := `Usage: %s COMMAND [OPTIONS] [ARGS...]

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

	return fmt.Sprintf(usage, self)
}

func Main(stdin io.Reader, stdout io.Writer, stderr io.Writer, self string, args ...string) error {
	if stdin == nil {
		return errors.New("no standard input")
	}

	if stdout == nil {
		return errors.New("no standard output")
	}

	if stderr == nil {
		return errors.New("no standard error")
	}

	if len(args) < 1 {
		help(self)
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if command := commands[arg]; command.main != nil {
			if err := command.main(args[i:]...); err != nil {
				return err
			}
		}

		switch arg {
		case "-c", "--config":
			if i++; i >= len(args) {
				return errors.New("config path not given")
			}

			conf = args[i]
		case "-v", "--verbose":
			verbose = true
		case "--version":
			fmt.Fprintln(stdout, version)
			return nil
		case "-h", "--help":
			fmt.Fprintln(stderr, help(self))
			return nil
		default:
			return errors.New("unknown command or flag: '" + arg + "'")
		}
	}

	return nil
}
