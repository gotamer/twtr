package quickstart

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
