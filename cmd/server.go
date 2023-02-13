package cmd

import (
	"fmt"

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
	err := s.Run(":8001")

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
