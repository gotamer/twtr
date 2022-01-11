package cmd

import (
	"errors"
	"fmt"
	"os"
)

const version = "v0.0.0"

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
	var cmd command

	if ctx == nil {
		ctx = &Context{}
	}

	if ctx.Self == "" {
		ctx.Self = "twtr"
	}

	if ctx.Config == "" {
		ctx.Config = "~/.config/twtxt/config"
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
		fmt.Fprint(ctx.Stderr, help(ctx))
		return nil
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]

		switch arg {
		case "-c", "--config":
			if i++; i >= len(args) {
				return errors.New("config path not given")
			}

			ctx.Config = args[i]
		case "-v", "--verbose":
			ctx.Verbose = true
		case "--version":
			fmt.Fprintln(ctx.Stdout, version)
			return nil
		case "-h", "--help":
			var msg string

			if cmd.name == "" {
				msg = help(ctx)
			} else {
				msg = cmd.help(ctx)
			}

			fmt.Fprint(ctx.Stderr, msg)
			return nil
		default:
			if c, ok := commands[arg]; ok {
				cmd = c
			} else {
				return errors.New("unknown command or flag: '" + arg + "'")
			}
		}
	}

	return nil
}
