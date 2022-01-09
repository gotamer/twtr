package cmd

import "io"

type Context struct {
	Self    string
	Config  string
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	Verbose bool
}
