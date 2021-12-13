// A twtxt client written in Go, hopefully the Go compatibility promise will
// prevent the issues that the original https://github.com/buckket/twtxt is
// experiencing due to the Python team making breaking changes.
package main

import (
	"fmt"
	"os"
	"path"
)

var (
	// version is the current distribution version.
	version = "v0.0.0"

	// verbose defines the level of detail to display.
	verbose bool

	// self is the name of the executable, defaults to twtxt.
	self string = "twtxt"
)

// command contains the main function of a subcommand as well as it's metadata
// and usage.
type command struct {
	main func(...[]string) error
	help func()
	desc string
}

// commands that can be executed via "twtxt <cmd> <params...>", these are
// essentially named main functions.
var commands map[string]command = map[string]command{
	"follow":     {nil, nil, "Add a new source to your followings."},
	"following":  {nil, nil, "Return the list of sources you're following."},
	"quickstart": {nil, nil, "Quickstart wizard for setting up twtxt."},
	"timeline":   {nil, nil, "Retrieve your personal timeline."},
	"tweet":      {nil, nil, "Append a new tweet to your files."},
	"unfollow":   {nil, nil, "Remove an existing source from your list."},
}

func usage() {
	if _, err := fmt.Fprintf(os.Stderr,
		`Usage: %s [OPTIONS] COMMAND [ARGS...]

	Decentralized, minimalist microblogging service for hackers.

Options:
	%s

Commands:
	%s
`,
		self,
		"TODO: List options",
		"TODO: List commands",
	); err != nil {
		panic(err)
	}
}

// fatal writes the given error message to the standard error stream and bails
// out with exit code 1.
func fatal(args ...interface{}) {
	msg := "fatal error"

	switch len(args) {
	case 1:
		msg = fmt.Sprint(args[0])
	default:
		if format, ok := args[0].(string); ok {
			msg = fmt.Sprintf(format, args[1:]...)
		} else {
			msg = fmt.Sprint(args...)
		}
	}

	if _, err := fmt.Fprintf(os.Stderr, "%s: %s\n", self, msg); err != nil {
		panic(fmt.Sprintf("fatal:\n%q\n\nwhile printing:\n%q", err, msg))
	}

	os.Exit(1)
}

// help is the main help and usage message, it exists the program with status
// code 0.
func help() {
	os.Exit(0)
}

// main is the entry point for the twtxt subcommands, it selects the given
// subcommand to run and executes it, or calls help() if the input was invalid.
func main() {
	// set the executable name (in case it was installed as something else)
	self = path.Base(os.Args[0])

	// select the subcommand (and or options)
	for i, arg := range os.Args[1:] {
		if cmd, ok := commands[arg]; ok {
			if cmd.main == nil {
				fatal("%s is not implemented yet", arg)
			}

			if err := cmd.main(os.Args[i:]); err != nil {
				fatal("%s: %s", arg, err)
			}

			os.Exit(0)
		}

		switch arg {
		case "-c", "--config":
			fatal("config parsing is not implemented yet")
		case "-v", "--verbose":
			verbose = true
		case "-V", "--version":
			fmt.Println(version)
			os.Exit(0)
		case "-h", "--help":
			help()
		default:
			fatal("unknown option: %s", arg)
		}
	}

	// show help message if no valid subcommand was given
	help()
}
