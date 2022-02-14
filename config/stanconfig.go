package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

type Nats struct {
	ClusterID string `yaml:"cluster_id"`
	ClientID  string `yaml:"client_id"`
	Channel   string `yaml:"channel"`
}

type NatsConfig struct {
	Nats Nats `yaml:"nats"`
}

func GetNatsConfigurations(mainConfigFile string) (cfg *NatsConfig) {
	file, err := os.Open(filepath.Clean(mainConfigFile))

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
