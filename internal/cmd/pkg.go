// Package cmd is an internal package that registers and executes subcommands of
// the twtr tool.
package cmd

import (
	"errors"
	"fmt"
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

// command represents a main function, i.e. a command, that can be called via
// the twtr tool.
type command interface {
	Main(...string) error
	Help() string
	Desc() string
}

var (
	// commands are the collection of registered commands, these are modified by
	// the Register function.
	commands map[string]command = map[string]command{
		"config":     config.Command,
		"follow":     follow.Command,
		"following":  following.Command,
		"quickstart": quickstart.Command,
		"timeline":   timeline.Command,
		"tweet":      tweet.Command,
		"unfollow":   unfollow.Command,
		"view":       view.Command,
	}

	// ErrUnknownCommand is returned when no command has been registered under
	// the requested name.
	ErrUnknownCommand = errors.New("command not found")
)

// Run looks for a registered command with the given name and runs it, if no
// command is found with that name, ErrUnknownCommand is returned.
func Run(name string, args ...string) error {
	if command, ok := commands[name]; ok {
		return command.Main(args...)
	}

	return ErrUnknownCommand
}

// Help generates a help message for all the subcommands as they are defined
// within their respective packages.
func Help(self string) {
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
