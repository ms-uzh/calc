package config

import (
	"github.com/fforootd/calc/models"

	"github.com/caos/utils/logging"

	"github.com/caos/utils/config/yaml"
)

type Config struct {
	Heads       models.Heads
	Polynamines models.Polyamines
	Tails       models.Tails
	Spiders     models.Spiders
	App         App
}

type App struct {
	MaxPolyamineSelectors uint
	URL                   string
	ServerPort            string
	RedirectPort          string
}

func ReadConfigs(headPath, polyaminePath, tailPath, spiederPath, appConfigPath string) *Config {
	config := new(Config)

	readObject(headPath, &config.Heads)
	readObject(polyaminePath, &config.Polynamines)
	readObject(tailPath, &config.Tails)
	readObject(appConfigPath, &config.App)
	readObject(spiederPath, &config.Spiders)

	if len(config.Heads) == 0 || len(config.Polynamines) == 0 || len(config.Tails) == 0 {
		logging.Log("CONF-1NL9X").
			WithField("heads", len(config.Heads)).
			WithField("polyamines", len(config.Polynamines)).
			WithField("tail", len(config.Tails)).
			WithField("spider", len(config.Spiders)).
			Panic("no objects configured")
	}

	return config
}

func readObject(path string, object interface{}) {
	err := yaml.ReadConfig(object, path)
	logging.Log("CONF-aNzV4").OnError(err).WithField("path", path).Panic("read configuration failed")
}
