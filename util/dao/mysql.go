package dao

import (
	"fmt"
	"follower/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var dbMap map[string]*gorm.DB

func Startup() {
	dbMap = make(map[string]*gorm.DB, len(config.App.Database))
	for _, databaseConfig := range config.App.Database {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			databaseConfig.Username, databaseConfig.Password, databaseConfig.Host, databaseConfig.Port, databaseConfig.Database, databaseConfig.Charset)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalln("connect " + dsn + " err")
		}

		dbMap[databaseConfig.Database] = db
	}
}

func GetMysqlDAO(key string) *gorm.DB {
	return dbMap[key]
}
