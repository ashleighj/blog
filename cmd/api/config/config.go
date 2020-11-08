package config

import (
	lg "api/logger"
	"context"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	defaultFilePath = "config/config.yml"
)

type Config struct {
	Service  Service    `json:"service"`
	Database DataSource `yaml:"database"`
	Redis    DataSource `yaml:"redis"`
}

type Service struct {
	Port  string `yaml:"port"`
	Debug bool   `yaml:"debug"`
}

type DataSource struct {
	Engine   string `yaml:"engine"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// GetDefault returns service config struct populated with
// values specified in a file in a default location
func GetDefault(ctx context.Context) (*Config, error) {
	return Get(ctx, defaultFilePath)
}

// Get returns service configuration struct populated with
// values specified in a file with the provided filename
func Get(ctx context.Context, filePath string) (conf *Config, e error) {
	var confFile []byte

	confFile, e = ioutil.ReadFile(filePath)
	if e != nil {
		lg.Error(ctx, e)
		return
	}

	if e == nil {
		e = yaml.Unmarshal(confFile, &conf)
		if e != nil {
			lg.Error(ctx, e)
			return
		}
	}

	return
}
