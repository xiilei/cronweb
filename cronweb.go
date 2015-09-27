package main

import (
	"github.com/codegangsta/cli"
	"github.com/xiilei/cronweb/commands"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "cronweb"
	app.Usage = "the linux crontab monitor"
	app.Version = "0.0.0"
	app.Commands = []cli.Command{
		commands.WebCmd,
		commands.RunCmd,
	}
	app.Run(os.Args)
}
