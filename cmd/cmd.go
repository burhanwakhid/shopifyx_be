package cmd

import (
	"os"

	"github.com/burhanwakhid/shopifyx_backend/config"
	"github.com/burhanwakhid/shopifyx_backend/internal/app/restserver"
	"github.com/urfave/cli"
)

func Initialize() {
	cfg := config.Initialize()
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:        "start-rest-server",
			Description: "Start REST server",
			Action:      startRESTserver(cfg),
		},
		{
			Name:        "",
			Description: "Start GRPC server",
			Action:      startRESTserver(cfg),
		},
	}

	app.Action = startRESTserver(cfg)

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func startRESTserver(cfg *config.Config) func(*cli.Context) error {
	return func(context *cli.Context) error {
		server := restserver.InitServer(cfg)
		return server.Start()
	}
}
