package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ke-db-migration/config"
	"log"
)

var (
	DB       *gorm.DB
	dbConfig *gorm.Config
)

func InitDb() {
	var err error
	dbConfig = &gorm.Config{}
	DB, err = gorm.Open(getConnection(config.Config.Database), dbConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func getConnection(conf config.Database) gorm.Dialector {
	if conf.Driver == "postgres" {
		return postgres.New(postgres.Config{DSN: fmt.Sprintf("host=%s user=%s "+
			"password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			conf.Host, conf.Username, conf.Password, conf.Database, conf.Port)})
	}
	return mysql.New(mysql.Config{DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charge=utf8&parseTime=True&loc=Local",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)})
}
