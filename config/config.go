package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Username string
	Password string
}

func ReadConfig(filename string) (*Config, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(raw, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
