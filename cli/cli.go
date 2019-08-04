package cli

import (
	"os"

	"github.com/urfave/cli"
)

type Application struct {
	*cli.App
	verbose    bool
	configFile string
}

func New() *Application {
	app := &Application{App: cli.NewApp()}
	app.Name = "rp"
	app.Version = "0.0.1"
	app.Usage = "reverse-proxy with weighted round-robin load-balancer."
	app.Author = "Ahmadreza Zibaei (ahmdrz)"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			EnvVar:      "VERBOSE",
			Name:        "verbose",
			Usage:       "verbose log output",
			Destination: &app.verbose,
		},
		cli.StringFlag{
			Name:        "config, c",
			Usage:       "load configuration from `FILE`",
			Destination: &app.configFile,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "Start reverse-proxy to serve on listen address",
			Action:  app.serve,
		},
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "Generate a simple configuration file from template",
			Action:  app.generate,
		},
	}

	return app
}

func (a *Application) Run() error {
	return a.App.Run(os.Args)
}
