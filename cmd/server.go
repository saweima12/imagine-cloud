package cmd

import (
	"fmt"

	"github.com/saweima12/imagine/internal/server"
	"github.com/urfave/cli/v2"
)

var CmdServer = cli.Command{
	Name:   "server",
	Action: runServer,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Value:   "8001",
			Usage:   "Temporary port number",
		},
	},
}

func runServer(c *cli.Context) error {
	s := server.New()

	port := c.String("port")
	portStr := fmt.Sprintf(":%v", port)

	err := s.Run(portStr)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
