package main

import (
	"os"

	cmdcmd "github.com/forbole/bookkeeper/cmd/cmd"
)

func main() {
	cmd := cmdcmd.ParseCmd()
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
