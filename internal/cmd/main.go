package cmd

import (
	"errors"
	"fmt"
	"os"
)

const version = "v0.0.0"

func Main(ctx *Context, args ...string) error {
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
			fmt.Fprint(ctx.Stderr, help(ctx))
			return nil
		default:
			return errors.New("unknown command or flag: '" + arg + "'")
		}
	}

	return nil
}
