package cmd

import (
	"errors"
	"fmt"
	"io"
)

const version = "v0.0.0"

var (
	conf    string = "~/.config/twtxt/config"
	verbose bool
)

func help(self string) string {
	usage := `Usage: %s COMMAND [OPTIONS] [ARGS...]

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
			fmt.Fprint(stderr, help(self))
			return nil
		default:
			return errors.New("unknown command or flag: '" + arg + "'")
		}
	}

	return nil
}
