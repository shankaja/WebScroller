package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Service Files `yaml:"service"`
}

type Files struct {
	Port string `yaml:"port"`
	URL  string `url:"resume"`
}

func LoadAll(configPath string) *Config {

	var config Config = Config{}

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &config
}
