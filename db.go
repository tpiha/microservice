package main

import (
	"flag"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initDb() *gorm.DB {

	var db *gorm.DB
	var err error

	if flag.Lookup("test.v") == nil {
		if db, err = gorm.Open(
			mysql.Open("root:root@tcp(database:3306)/microservice?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Warn),
			}); err != nil {

			log.Printf("[initDb] error: %s", err)

			db, _ = gorm.Open(sqlite.Open("microservice.db"), &gorm.Config{})

		}
	} else {
		db, _ = gorm.Open(sqlite.Open("microservice.db"), &gorm.Config{})
	}

	if err := db.AutoMigrate(&Metric{}, &Datapoint{}); err != nil {
		panic(err.Error())
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	return db
}
