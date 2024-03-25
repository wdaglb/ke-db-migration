package core

import (
	"fmt"
	gologger "github.com/phachon/go-logger"
	"ke-db-migration/config"
	"os"
)

var (
	Logger *gologger.Logger
)

func InitApp() {
	_ = os.MkdirAll(config.Config.LogDir, os.ModePerm)
	Logger = gologger.NewLogger()
	_ = Logger.Detach("console")
	_ = Logger.Attach("console", gologger.LOGGER_LEVEL_DEBUG, &gologger.ConsoleConfig{
		Color:      true,
		JsonFormat: false,
	})
	_ = Logger.Attach("file", gologger.LOGGER_LEVEL_DEBUG, &gologger.FileConfig{
		Filename:   fmt.Sprintf("%s/migrate.log", config.Config.LogDir),
		MaxSize:    1024 * 1024,
		MaxLine:    100000,
		DateSlice:  "d",
		JsonFormat: false,
		Format:     "",
	})
	initDb()
}
