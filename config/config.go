package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var Config config

type config struct {
	Database     Database
	MigrationDir string `yaml:"migrationDir"`
	MigrationDb  string `yaml:"migrationDb"`
}

func init() {
	dataBytes, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(dataBytes, &Config)
	if err != nil {
		log.Fatal(err)
	}
}
