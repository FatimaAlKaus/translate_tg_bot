package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	API_KEY       string `yaml:"tg_api_key"`
	RAPID_API_KEY string `yaml:"rapid_api_key"`
}

func Load(file string) (*Config, error) {
	var config Config

	// load from YAML config file
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
