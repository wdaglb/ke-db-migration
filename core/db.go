package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"ke-db-migration/config"
)

var (
	DB       *gorm.DB
	dbConfig *gorm.Config
)

func initDb() {
	var err error
	dbConfig = &gorm.Config{}
	DB, err = gorm.Open(getConnection(config.Config.Database), dbConfig)
	if err != nil {
		Logger.Errorf(err.Error())
	}
	DB.Logger.LogMode(logger.Error)
}

func getConnection(conf config.Database) gorm.Dialector {
	if conf.Driver == "postgres" {
		return postgres.New(postgres.Config{DSN: fmt.Sprintf("host=%s user=%s "+
			"password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			conf.Host, conf.Username, conf.Password, conf.Database, conf.Port)})
	}
	return mysql.New(mysql.Config{DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.Database, conf.Charset)})
}
