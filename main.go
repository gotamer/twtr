package main

import (
	"fmt"
	"os"
	"path"

	"internal/cmd"
)

func main() {
	ctx := cmd.Context{
		Self:   path.Base(os.Args[0]),
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if err := cmd.Main(&ctx, os.Args[1:]...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
