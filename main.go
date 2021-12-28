package main

import (
	"fmt"
	"os"
	"path"
)

var (
	self    string = "twtr"
	conf    string
	verbose bool
	version string = "v0.0.0"
)

var commands map[string]func(...string) error = map[string]func(...string) error{
	"config": main_config,
}

func help() {
	usage := `Usage: %s COMMAND [OPTIONS] [ARGS...]

Decentralized, minimalist microblogging for hackers.

Options:
    -c, --config PATH  Specify a custom config file location.
    -v, --version      Enable verbose output for debugging.
    --version          Show the version and exit.
    -h, --help         Show this message and exit.

Command:
    timeline   Retrieve your personal timeline.
    following  View the sources that you are following.
    follow     Add a new source to your followings.
    unfollow   Remove an existing source from your list.
    tweet      Send out a message into the void.
    view       View a source that you follow.
    config     Update your configuration.
`

	fmt.Fprintf(os.Stderr, usage, self)
}

func main() {
	self = path.Base(os.Args[0])
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if cmd, ok := commands[arg]; ok {
			if err := cmd(args[i:]...); err != nil {
				panic(err)
			}

			os.Exit(0)
		}

		switch arg {
		case "-c", "--config":
			if i++; i >= len(args) {
				panic("config path not given")
			}

			conf = args[i]
		case "-v", "--verbose":
			verbose = true
		case "--version":
			fmt.Println(version)
			os.Exit(0)
		case "-h", "--help":
			help()
			os.Exit(0)
		default:
			panic("unknown command or flag: '" + arg + "'")
		}
	}
}
