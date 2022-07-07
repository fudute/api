package config

import (
	"bytes"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Template   string `yaml:"template"`
	EnableACL  bool   `yaml:"enable_acl"`
	EnableDiff bool   `yaml:"enable_diff"`
	Image      string `yaml:"image"`
}

var Conf = new(Config)

func loadConfig() {
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	dec := yaml.NewDecoder(bytes.NewReader(configFile))
	if err := dec.Decode(Conf); err != nil {
		log.Fatal(err)
	}
	log.Printf("config is: %+v", Conf)
}
