package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config - structure for reading the configuration file.
type Config struct {
	Listen string `yaml:"listen"`
}

func ReadConfig(ConfigName string) (x *Config, err error) {
	var file []byte
	if file, err = ioutil.ReadFile(ConfigName); err != nil {
		return nil, err
	}
	x = new(Config)
	if err = yaml.Unmarshal(file, x); err != nil {
		return nil, err
	}
	return x, nil
}
