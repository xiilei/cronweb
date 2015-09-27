package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
)

var RunCmd = cli.Command{
	Name:   "run",
	Usage:  "the agent of task",
	Action: runTask,
}

func runTask(c *cli.Context) {
	name := "nothing"
	if len(c.Args()) > 0 {
		name = c.Args()[0]
	}
	fmt.Println(name)
}
