package cmd

import (
	"hanmantpatil/gorilla/apis"
	"hanmantpatil/gorilla/config"
	"hanmantpatil/gorilla/service/db"
	"hanmantpatil/gorilla/service/logger"

	"github.com/urfave/cli"
	"go.uber.org/fx"
)

var gorillaCmd = &cli.Command{
	Name: "gorilla",
	Action: func(ctx *cli.Context) error {
		fx.New(
			fx.Provide(config.GetConfig),
			fx.Provide(db.GetInstance),
			fx.Provide(logger.GetInstance),

			repositories(),
			usecases(),
			fx.Provide(apis.NewHander),

			fx.Invoke(func(handler *apis.Handler) {
				handler.RegisterRoutes().Start("localhost:8080")
			}),
		).Run()

		return nil
	},
}
