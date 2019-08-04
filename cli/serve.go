package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	rp "github.com/ahmdrz/rp/reverse-proxy"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

type target struct {
	Address string
	Weight  int
}

type config struct {
	ListenAddr string
	Targets    []target
}

func (c *config) Save(path string) error {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bytes, 0755)
}

func loadConfig(path string) (*config, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &config{}
	return cfg, yaml.Unmarshal(bytes, cfg)
}

func (a *Application) serve(c *cli.Context) error {
	if a.configFile == "" {
		a.configFile = "rpconfig.yaml"
	}
	_, err := os.Stat(a.configFile)
	if os.IsNotExist(err) {
		return fmt.Errorf("config file [%s] : no such file or directory", a.configFile)
	}

	cfg, err := loadConfig(a.configFile)
	if err != nil {
		return err
	}

	proxy := rp.New()
	proxy.Log(a.verbose)

	for _, item := range cfg.Targets {
		targetURL, err := url.Parse(item.Address)
		if err != nil {
			return err
		}
		proxy.Add(targetURL, item.Weight)
	}

	log.Printf("Starting reverse-proxy server on %s ...", cfg.ListenAddr)
	return proxy.ListenAndServe(cfg.ListenAddr)
}
