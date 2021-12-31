package main

import (
	"fmt"
	"os"
	"path"

	"internal/cmd"
)

const version = "v0.0.0"

var (
	conf    string
	verbose bool
)

func main() {
	self := path.Base(os.Args[0])
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		arg := args[i]

		switch err := cmd.Run(arg, args[i:]...); err {
		case nil:
			os.Exit(0)
		case cmd.ErrUnknownCommand:
			break
		default:
			panic(err)
		}

		switch arg {
		case "-c", "--config":
			if i++; i >= len(args) {
				panic("config path not given")
			}

			conf = args[i]
		case "-v", "--verbose":
			verbose = true
		case "--version":
			fmt.Println(version)
			os.Exit(0)
		case "-h", "--help":
			cmd.Help(self)
			os.Exit(0)
		default:
			panic("unknown command or flag: '" + arg + "'")
		}
	}
}
