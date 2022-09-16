package cmd

import (
	"github.com/saweima12/imagine/internal/imagine"
	"github.com/urfave/cli/v2"
)

var CmdServer = cli.Command{
	Name:   "server",
	Action: runServer,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Value:   "8000",
			Usage:   "Temporary port number",
		},
	},
}

func runServer(c *cli.Context) error {
	s := imagine.New()
	s.Run(":8001")
	return nil
}
