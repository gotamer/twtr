// Package cmd is an internal package that registers and executes subcommands of
// the twtr tool.
package cmd

import "errors"

// command represents a main function, i.e. a command, that can be called via
// the twtr tool.
type command func(...string) error

// commands are the collection of registered commands, these are modified by the
// Register function.
var commands map[string]command = make(map[string]command)

var (
	// ErrDuplicateCommand is returned when a command cannot be registered
	// because a command with that name has already been registered.
	ErrDuplicateCommand = errors.New("command already registered")

	// ErrUnknownCommand is returned when no command has been registered under
	// the requested name.
	ErrUnknownCommand = errors.New("command not found")
)

// Register registers a named command to the commands list, returns
// ErrDuplicateCommand if the given name already exists.
func Register(name string, main command) error {
	if _, ok := commands[name]; !ok {
		commands[name] = main

		return nil
	}

	return ErrDuplicateCommand
}

// Run looks for a registered command with the given name and runs it, if no
// command is found with that name, ErrUnknownCommand is returned.
func Run(name string, args ...string) error {
	if main, ok := commands[name]; ok && main != nil {
		return main(args...)
	}

	return ErrUnknownCommand
}
