package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// ServerConfig ...
type ServerConfig struct {
	Port string `yaml:"port"`
}

// Config ...
type Config struct {
	Server ServerConfig `yaml:"server"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	configBytes, err := ioutil.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	configBytes = []byte(os.ExpandEnv(string(configBytes)))
	var config Config
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
