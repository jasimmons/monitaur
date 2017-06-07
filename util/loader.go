package util

import (
	"encoding/json"
	"errors"
	"github.com/jasimmons/monitaur"
	"github.com/jasimmons/monitaur/log"
	"io/ioutil"
	"os"
	"path/filepath"
)

func LoadConfig(configDir string) ([]monitaur.Check, []monitaur.Handler, error) {
	if configDir == "" {
		log.Info("no configuration directory specified: not loading configuration")
		return nil, nil, nil
	}

	checks := make([]monitaur.Check, 0)
	handlers := make([]monitaur.Handler, 0)
	err := filepath.Walk(configDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			log.Infof("ignoring directory: %s\n", path)
			return nil
		} else {

			if filepath.Ext(path) != ".json" {
				log.Infof("file at %s has extension '%s', not '.json'\n", path, filepath.Ext(path))
				return nil
			}
			var fContent []byte
			if fContent, err = ioutil.ReadFile(path); err != nil {
				return err
			}

			if len(fContent) == 0 {
				log.Infof("read empty JSON file -- ignoring (at: %s)\n", path)
				return nil
			}

			type cfgType struct {
				Type string `json:"type"`
			}
			var cfg cfgType
			err := json.Unmarshal(fContent, &cfg)
			if err != nil {
				return err
			}

			switch cfg.Type {
			case monitaur.TYPE_CHECK:
				// load into a check struct
				// append check to slice
				check, err := monitaur.ParseCheckWithDefaults(fContent)
				if err != nil {
					return err
				}
				checks = append(checks, check)
			case monitaur.TYPE_HANDLER:
				// load into a handler struct
				// append handler to slice
				handler, err := monitaur.ParseHandlerWithDefaults(fContent)
				if err != nil {
					return err
				}
				handlers = append(handlers, handler)
			default:
				// error with invalid type
				return errors.New(cfg.Type + " is not a valid configuration type")
			}
		}
		return nil
	})

	return nil, nil, err
}
