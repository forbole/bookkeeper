package main
import (
	cmdcmd "github.com/forbole/bookkeeper/cmd/cmd"
	"os"


)
func main() {
	cmd := cmdcmd.ParseCmd()
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
