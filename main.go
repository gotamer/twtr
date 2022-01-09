package main

import (
	"fmt"
	"os"
	"path"

	"internal/cmd"
)

func main() {
	stdin, stdout, stderr := os.Stdin, os.Stdout, os.Stderr
	self, args := path.Base(os.Args[0]), os.Args[1:]

	if err := cmd.Main(stdin, stdout, stderr, self, args...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
