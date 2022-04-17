package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/VxVxN/mdserver/pkg/consts"
)

var Cfg *config

// config - structure for reading the configuration file.
type config struct {
	Listen   string `yaml:"listen"`
	IsSSL    bool   `yaml:"ssl"`
	LevelLog LVLLog `yaml:"level_log"`
	// SessionAge - in minutes
	SessionAge    int    `yaml:"session_age"`
	SessionSecret string `yaml:"session_secret"`
	Domain        string `yaml:"domain"`
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

func InitTestConfig() {
	Cfg = &config{
		Listen:        "",
		LevelLog:      "",
		SessionAge:    120,
		SessionSecret: "secret",
		Domain:        "testDomain.com",
	}
}

func (cfg config) GetURL() string {
	protocol := consts.Http
	if cfg.IsSSL {
		protocol = consts.Https
	}
	return protocol + cfg.Domain
}
