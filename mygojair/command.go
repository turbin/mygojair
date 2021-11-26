package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var command_run = cli.Command{
	Name:      "run",
	ShortName: "",
	Aliases:   []string{},
	Usage: `Create a jair to isolate the process running
			in a isolated enviroment`,

	Action: func(context *cli.Context) error {

		// if len(context.Args()) < 1 {
		// 	return fmt.Errorf("Command No Found")
		// }

		var cmdArry []string
		for _, arg := range context.Args() {
			cmdArry = append(cmdArry, arg)
		}

		fmt.Fprintln(os.Stdout, "cmdArray:", cmdArry)

		//_ := context.Bool("it")
		Run(true, cmdArry)
		return nil
	},
}

var command_init = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		log.Infof("init come on")
		err := RunContainerInitProcess()
		if err != nil {
			log.Info("err %v", err)
		}
		return nil
	},
}
