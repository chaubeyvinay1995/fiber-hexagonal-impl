package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"os"
)

// Config struct
type Config struct {
	Database struct {
		URI     string `yaml:"uri"`
		URL     string `yaml:"url"`
		DB      string `yaml:"db"`
		Timeout int    `yaml:"timeout"`
	} `yaml:"database"`
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

func NewConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, errors.New("config file missing")
	}
	defer file.Close()
	cfg := &Config{}
	yd := yaml.NewDecoder(file)
	err = yd.Decode(cfg)
	if err != nil {
		return nil, errors.New("error while decoding")
	}
	return cfg, nil
}
