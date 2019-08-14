package commands

import (
	"github.com/urfave/cli"
)

// Serve - For serve the application
var Serve = cli.Command{
	Name:        "serve",
	Usage:       "Run service HTTP Server for serving REST API",
	Description: "Execute command will start the REST API",
	Action:      ServeHTTP,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "./App.yaml",
			Usage: "Directory configuration",
		},
	},
}
