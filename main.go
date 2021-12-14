// A twtxt client written in Go, hopefully the Go compatibility promise will
// prevent the issues that the original https://github.com/buckket/twtxt is
// experiencing due to the Python team making breaking changes.
package main

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
)

var (
	// version is the current distribution version.
	version = "v0.0.0"

	// verbose defines the level of detail to display.
	verbose bool

	// self is the name of the executable, defaults to twtxt.
	self string = "twtxt"

	// conf is the location of the config, defaults to ~/.config/twtxt/config
	conf string
)

// command contains the main function of a subcommand as well as it's metadata
// and usage.
type command struct {
	main func(...[]string) error
	desc string
}

// commands that can be executed via "twtxt <cmd> <params...>", these are
// essentially named main functions.
var commands map[string]command = map[string]command{
	"follow":     {nil, "Add a new source to your followings."},
	"following":  {nil, "Return the list of sources you're following."},
	"quickstart": {nil, "Quickstart wizard for setting up twtxt."},
	"timeline":   {nil, "Retrieve your personal timeline."},
	"tweet":      {nil, "Append a new tweet to your files."},
	"unfollow":   {nil, "Remove an existing source from your list."},
}

func printErr(a ...interface{}) {
	if _, err := fmt.Fprint(os.Stderr, a...); err != nil {
		panic(err)
	}
}

func printErrf(format string, a ...interface{}) {
	if _, err := fmt.Fprintf(os.Stderr, format, a...); err != nil {
		panic(err)
	}
}

func printErrln(a ...interface{}) {
	if _, err := fmt.Fprintln(os.Stderr, a...); err != nil {
		panic(err)
	}
}

// help is the main help and usage message, it exists the program with status
// code 0.
func help(msg string) {
	if msg == "" {
		defer os.Exit(0)
	} else {
		defer os.Exit(1)

		printErrln(msg)
	}

	cmds := make([]string, len(commands), len(commands))

	i := 0
	for cmd := range commands {
		cmds[i] = cmd
		i++
	}

	sort.Strings(cmds)

	for i, cmd := range cmds {
		cmds[i] = fmt.Sprintf("%-16s%s", cmd, commands[cmd].desc)
	}

	printErrf(`Usage: %s [OPTIONS] COMMAND [ARGS...]

	Decentralized, minimalist microblogging service for hackers.

Options:
	%s

Commands:
	%s

version %s - all rights reversed

`,
		self,
		"TODO: List options",
		strings.Join(cmds, "\n\t"),
		version,
	)
}

// main is the entry point for the twtxt subcommands, it selects the given
// subcommand to run and executes it, or calls help() if the input was invalid.
func main() {
	var skip int

	// set the executable name (in case it was installed as something else)
	self = path.Base(os.Args[0])

	// select the subcommand (and or options)
	for i, arg := range os.Args[1:] {
		if skip > 0 {
			skip--
			continue
		}

		if cmd, ok := commands[arg]; ok {
			if cmd.main == nil {
				printErrln(arg + " is not implemented yet")
			}

			if err := cmd.main(os.Args[i:]); err != nil {
				printErrln(arg + ": " + err.Error())
			}

			os.Exit(0)
		}

		switch arg {
		case "-c", "--config":
			if len(os.Args[i:]) > 0 {
				conf = os.Args[i+1]
				skip++
			} else {
				help(fmt.Sprintf("option '%s' requires PATH", arg))
			}
		case "-v", "--verbose":
			verbose = true
		case "-V", "--version":
			fmt.Println(version)
			os.Exit(0)
		case "-h", "--help":
			help("")
		default:
			help(fmt.Sprintf("unknown option: %s", arg))
		}
	}

	// show help message if no valid subcommand was given
	help("no subcommand given")
}
