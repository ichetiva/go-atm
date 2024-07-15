package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Addr         string `yaml:"addr"`
	LoggingLevel string `yaml:"logging_level"`
}

func LoadConfig(path string) (Config, error) {
	var config Config
	contents, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(contents, &config)
	return config, err
}
