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

// PostgresConfig ...
type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
	SSLMode  string `yaml:"ssl_mode"`
}

// Config ...
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Postgres PostgresConfig `yaml:"postgres"`
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
