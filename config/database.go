package config

type Database struct {
	Driver   string
	Host     string
	Port     uint
	Database string
	Schema   string
	Username string
	Password string
}
