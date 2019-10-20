package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	JWT struct {
		Secret    string `yaml:"secret"`
	} `yaml:"jwt"`
}

func GetConfig() Config {
	f, err := os.Open("./../configs/config.yaml")
	if err != nil {
		processError(err)
	}

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		processError(err)
	}
	return cfg
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
