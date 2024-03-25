package core

import (
	gologger "github.com/phachon/go-logger"
)

var (
	Logger *gologger.Logger
)

func InitApp() {
	Logger = gologger.NewLogger()
	_ = Logger.Attach("console", gologger.LOGGER_LEVEL_DEBUG, &gologger.ConsoleConfig{
		Color:      true,
		JsonFormat: false,
	})
	_ = Logger.Attach("file", gologger.LOGGER_LEVEL_DEBUG, &gologger.FileConfig{
		Filename:   "sql.log",
		JsonFormat: true,
		Format:     "",
	})
	initDb()
}
