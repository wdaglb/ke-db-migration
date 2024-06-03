package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var Config config

type config struct {
	Database     Database
	Notify       Notify `yaml:"notify"`
	MigrationDir string `yaml:"migrationDir"`
	MigrationDb  string `yaml:"migrationDb"`
	LogDir       string `yaml:"logDir"`
	EnableLog    bool   `yaml:"enableLog"`
}

func init() {
	dir := flag.String("config", "config.yml", "配置文件")

	dataBytes, err := os.ReadFile(*dir)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(dataBytes, &Config)
	if err != nil {
		log.Fatal(err)
	}
}
