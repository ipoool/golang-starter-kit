package main

import (
	"log"
	"os"

	"github.com/ipoool/laporcuranmor-api/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Golang Starter Kit"
	app.Usage = "Golang Starter Kit - Lets build your awesome application"
	app.Version = "1.0.0"
	app.Commands = cli.Commands{
		commands.Serve,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
