package main

import (
	"os"

	"github.com/callicoder/go-ready/cmd"
)

func main() {
	cli := cmd.NewCLI()
	cli.SetArgs(os.Args[1:])

	err := cli.Execute()
	if err != nil {
		os.Exit(1)
	}
}
