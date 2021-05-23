package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var Cfg *config

// config - structure for reading the configuration file.
type config struct {
	Listen   string `yaml:"listen"`
	LevelLog LVLLog `yaml:"level_log"`
}

type LVLLog string

const (
	CommonLog LVLLog = "common"
	DebugLog         = "debug"
	TraceLog         = "trace"
)

func InitConfig(ConfigName string) (err error) {
	var file []byte
	if file, err = ioutil.ReadFile(ConfigName); err != nil {
		return err
	}
	Cfg = &config{}
	if err = yaml.Unmarshal(file, Cfg); err != nil {
		return err
	}
	return nil
}
