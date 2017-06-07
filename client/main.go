package main

import (
	"flag"
	"github.com/jasimmons/monitaur"
	"github.com/jasimmons/monitaur/log"
	"github.com/jasimmons/monitaur/util"
)

var configDir string

func init() {
	const (
		defaultConfigDir = "/etc/monitaur/conf.d"
		usage            = "the directory where configuration is located"
	)
	flag.StringVar(&configDir, "c", defaultConfigDir, usage)
}

func main() {
	log.Info("monitaur!")
	flag.Parse()

	cfg := &monitaur.Config{
		ConfigDir: configDir,
	}
	checks, _, err := util.LoadConfig(cfg.ConfigDir)
	if err != nil {
		log.Fatal(err)
	}

	client := monitaur.NewClient(cfg, checks)
	client.Run()
}
