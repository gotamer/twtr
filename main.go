package main

import (
	"fmt"
	"os"

	"duriny.envs.sh/twtr/internal/cmd"
)

func main() {
	if err := cmd.Main(nil, os.Args[1:]...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
