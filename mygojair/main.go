package main

import (
	"log"
	"os"

	// log "github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

const Usage = `my go jair is my first docker like app.`

func main() {
	app := cli.NewApp()
	app.Name = "mygojair"
	app.Usage = Usage

	app.Commands = []cli.Command{
		command_run,
		command_init,
	}

	app.Before = func(c *cli.Context) error {
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
