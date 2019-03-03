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

	err := readObject(headPath, &config.Heads)
	logging.Log("CONF-aNzV4").OnError(err).WithField("path", headPath).Panic("read head configuration failed")
	err = readObject(polyaminePath, &config.Polynamines)
	logging.Log("CONF-c9weA").OnError(err).WithField("path", polyaminePath).Panic("read polyamine configuration failed")
	err = readObject(tailPath, &config.Tails)
	logging.Log("CONF-qVvHO").OnError(err).WithField("path", tailPath).Panic("read tail configuration failed")

	return config
}

func readObject(path string, object interface{}) error {
	return yaml.ReadConfig(object, path)
}
