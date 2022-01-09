package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const version = "v0.0.0"

type Context struct {
	Self   string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

var (
	conf    string = "~/.config/twtxt/config"
	verbose bool
)

func help(ctx *Context) string {
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

	return fmt.Sprintf(usage, ctx.Self)
}

func Main(ctx *Context, args ...string) error {
	if ctx.Self == "" {
		ctx.Self = "twtr"
	}

	if ctx.Stdin == nil {
		ctx.Stdin = os.Stdin
	}

	if ctx.Stdout == nil {
		ctx.Stdout = os.Stdout
	}

	if ctx.Stderr == nil {
		ctx.Stderr = os.Stderr
	}

	if len(args) < 1 {
		help(ctx)
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
			fmt.Fprintln(ctx.Stdout, version)
			return nil
		case "-h", "--help":
			fmt.Fprint(ctx.Stderr, help(ctx))
			return nil
		default:
			return errors.New("unknown command or flag: '" + arg + "'")
		}
	}

	return nil
}
