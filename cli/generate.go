package cli

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func (a *Application) generate(c *cli.Context) error {
	if a.configFile == "" {
		a.configFile = "rpconfig.yaml"
	}
	_, err := os.Stat(a.configFile)
	if !os.IsNotExist(err) {
		return fmt.Errorf("config file [%s] : already exists", a.configFile)
	}
	cfg := &config{
		ListenAddr: "0.0.0.0:8080",
		Targets: []target{
			target{
				Address: "http://gstatic.com",
				Weight:  1,
			},
		},
	}
	return cfg.Save(a.configFile)
}
