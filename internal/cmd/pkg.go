// Package cmd is an internal package that registers and executes subcommands of
// the twtr tool.
package cmd

import (
	"errors"

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
type command func(...string) error

// commands are the collection of registered commands, these are modified by the
// Register function.
var commands map[string]command = map[string]command{
	"config":     config.Main,
	"follow":     follow.Main,
	"following":  following.Main,
	"quickstart": quickstart.Main,
	"timeline":   timeline.Main,
	"tweet":      tweet.Main,
	"unfollow":   unfollow.Main,
	"view":       view.Main,
}

// ErrUnknownCommand is returned when no command has been registered under
// the requested name.
var ErrUnknownCommand = errors.New("command not found")

// Run looks for a registered command with the given name and runs it, if no
// command is found with that name, ErrUnknownCommand is returned.
func Run(name string, args ...string) error {
	if main, ok := commands[name]; ok && main != nil {
		return main(args...)
	}

	return ErrUnknownCommand
}
