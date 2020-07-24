package main

import (
	"fmt"
	"pqredis/cmd"
)

var (
	version = "0.0.1"
	commit  = "n/a"
)

func main() {
	cli := cmd.NewCLI()
	cli.Version = fmt.Sprintf("%s (Commit: %s)", version, commit)
	cli.Execute()
}
