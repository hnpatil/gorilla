package cmd

import (
	"fmt"
	"hanmantpatil/gorilla/repository"
	"hanmantpatil/gorilla/usecase"
	"os"
	"time"

	"github.com/urfave/cli"
	"go.uber.org/fx"
)

func Run() {
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

var app = &cli.App{
	Name:      "gorilla",
	Compiled:  time.Now(),
	Copyright: "(c) 2023 - now Herewith GmbH",
	Usage:     "services and utilities to run the Gorilla backend",
	Commands: []cli.Command{
		*gorillaCmd,
		*workerCmd,
	},
	ExitErrHandler: func(cCtx *cli.Context, err error) {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func repositories() fx.Option {
	return fx.Provide(
		repository.NewBooks,
		repository.NewUsers,
		repository.NewAuthCodes,
	)
}

func usecases() fx.Option {
	return fx.Provide(
		usecase.NewBooks,
		usecase.NewUsers,
		usecase.NewAuth,
	)
}
