package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Host      string `yaml:"host"`
	Port      uint16 `yaml:"port"`
	RecordAll bool   `yaml:"record_all"`
	ShowURL   string `yaml:"show_url"`
	ResetURL  string `yaml:"reset_url"`
}

func DefaultConfig() Config {
	return Config{
		Host:      "0.0.0.0",
		Port:      8080,
		RecordAll: true,
		ShowURL:   "showmewhatyougot",
		ResetURL:  "resetwhatyougot",
	}
}

func LoadConfig(path string) (Config, error) {
	yamlstr, err := ioutil.ReadFile(path)
	// Initialize with default values
	yamlparsed := DefaultConfig()
	if err != nil {
		return yamlparsed, err
	}

	err = yaml.Unmarshal(yamlstr, &yamlparsed)
	if err != nil {
		return yamlparsed, err
	}

	return yamlparsed, nil
}
