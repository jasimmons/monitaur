package main

import (
	"flag"
	"github.com/jasimmons/monitaur"
	"github.com/jasimmons/monitaur/log"
	"github.com/jasimmons/monitaur/util"
)

var (
	configDir  string
	logVerbose bool
)

func init() {
	const (
		defaultConfigDir  = "/etc/monitaur/conf.d"
		configDirUsage    = "the directory where configuration is located"
		defaultLogVerbose = false
		logVerboseUsage   = "enable verbose logging (for debugging)"
	)
	flag.StringVar(&configDir, "c", defaultConfigDir, configDirUsage)
	flag.BoolVar(&logVerbose, "v", defaultLogVerbose, logVerboseUsage)
}

func main() {
	log.Info("monitaur!")
	flag.Parse()

	if logVerbose {
		log.Verbose = true
		log.Debug("debug logging enabled")
	}

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
