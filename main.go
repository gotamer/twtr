// NAME
//
// twtr - decentralized microblogging client
//
// SYNOPSIS
//
// twtr [OPTIONS] [COMMAND] [ARGS ...]
//
// DESCRIPTION
//
// twtr is a drop in replacement for the original twtxt client.
//
// You want to get some thoughts out on the internet in a convenient and
// slick way, while also following the gibberish of others? Instead of
// signing up to a closed source and regulated platform, getting your status
// updates out is as easy as adding a line to a publicly accessible text
// file. The URL pointing to this file is your identity, your account. twtr
// then tracks these text files, like a feedreader, and builds your unique
// timeline from the text files you follow. The format is simple, human
// readable, and integrates well with UNIX command line tools.
//
// tldr: twtr is a CLI tool for the twtxt self-hosted microblogging format.
//
// OPTIONS
//
// SUBCOMMANDS
//
// EXIT STATUS
//
// FILES
//
// ENVIRONMENT
//
// CONFORMING TO
//
// twtr conforms to the twtxt file specification, traditionally the file is
// located at https://example.com/path/to/twtxt.txt, however, as not
// everyone has access to a personal website to host their feeds, twtr also
// supports specialised hosting options, such as a GitHub gist.
//
// See https://twtxt.readthedocs.io/en/latest/user/twtxtfile.html for more
// information on the file structure.
//
// NOTES
//
// The original client was written is Python around 2016, and a small user
// base has been built around the twtxt format. Since the format is human
// readable and can be easily used with just shell commands, in addition to
// the original client, many users have written their own or just use the
// echo command.
//
// This client aims to be a complete drop-in replacement for the original
// client, not only to replicate the original feature set, but also to
// support many additions that the community of users have requested. There
// have also been a number of issues with the original client breaking
// because of backwards compatibility issues with the Python language. twtr
// aims to be a permanently supported tool, the Go language protects its
// backwards compatibility, so twtr will work forever!
//
// COPYRIGHT
//
// All rites reversed, use, distribute, and modify freely.
//
// AUTHOR
//
// ~duriny <duriny@envs.net>
//
// BUGS
//
// Probably. Let me know if you find any.
//
// SEE ALSO
//
// twtxt(1) - https://github.com/buckket/twtxt
//
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
	var status int

	if msg != "" {
		status = 1
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

	os.Exit(status)
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
