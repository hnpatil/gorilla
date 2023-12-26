package cmd

import (
	"time"

	"github.com/urfave/cli"
)

var workerCmd = &cli.Command{
	Name: "worker",
	Action: func(ctx *cli.Context) error {
		for true {
			println("working")

			time.Sleep(time.Second)
		}

		return nil
	},
}
