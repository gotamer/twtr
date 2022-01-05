package config

import (
	"errors"
)

type command struct{}

func (c command) Main(args ...string) error {
	return errors.New("unimplemented")
}

func (c command) Help() string {
	return ""
}

func (c command) Desc() string {
	return ""
}

var Command command = command{}

func get(key string) error {
	return errors.New("unimplemented")
}

func set(key, value string) error {
	return errors.New("unimplemented")
}

func remove(key string) error {
	return errors.New("unimplemented")
}

func edit() error {
	return errors.New("unimplemented")
}
