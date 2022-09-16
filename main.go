package main

import (
	"os"

	"github.com/saweima12/imagine/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:        "Imagine",
		Usage:       "Aw self-host Image storage service",
		Description: "By default, imgine will start serving using the webserver with no arguments.",
	}

	// register commnad
	app.Commands = []*cli.Command{
		&cmd.CmdServer,
	}

	// Run CLI Tools.
	app.Run(os.Args)
}
