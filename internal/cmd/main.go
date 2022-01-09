package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

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
	Self    string = "twtxt" // TODO: get from exec path
	Config  string = "~/.config/twtxt/config"
	Verbose bool
	Stdin   io.Reader = os.Stdin
	Stdout  io.Writer = os.Stdout
	Stderr  io.Writer = os.Stderr
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

func help() {
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

	fmt.Fprintf(Stderr, usage, Self)
}

func Main(args ...string) error {
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

			Config = args[i]
		case "-v", "--verbose":
			Verbose = true
		case "--version":
			fmt.Fprintln(Stdout, version)
			return nil
		case "-h", "--help":
			help()
			return nil
		default:
			return errors.New("unknown command or flag: '" + arg + "'")
		}
	}

	return nil
}
