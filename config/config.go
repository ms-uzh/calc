package config

import (
	"github.com/fforootd/calc/models"

	"github.com/yabslabs/utils/logging"

	"github.com/yabslabs/utils/config/yaml"
)

type Config struct {
	Heads       []models.Head
	Polynamines []models.Polyamine
	Tails       []models.Tail
}

func ReadConfigs(headPath, polyaminePath, tailPath string) *Config {
	config := new(Config)

	readObject(headPath, &config.Heads)
	readObject(polyaminePath, &config.Polynamines)
	readObject(tailPath, &config.Tails)

	if len(config.Heads) == 0 || len(config.Polynamines) == 0 || len(config.Tails) == 0 {
		logging.Log("CONF-1NL9X").
			WithField("heads", len(config.Heads)).
			WithField("polyamines", len(config.Polynamines)).
			WithField("tail", len(config.Tails)).
			Panic("no objects configured")
	}

	return config
}

func readObject(path string, object interface{}) {
	err := yaml.ReadConfig(object, path)
	logging.Log("CONF-aNzV4").OnError(err).WithField("path", path).Panic("read configuration failed")
}
