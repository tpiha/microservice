package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initDb() *gorm.DB {

	var db *gorm.DB
	var err error

	if db, err = gorm.Open(
		mysql.Open("root:root@tcp(database:3306)/microservice?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
			Logger:          logger.Default.LogMode(logger.Error),
			CreateBatchSize: 100,
		}); err != nil {

		log.Printf("[initDb] error: %s", err)
	}

	if err := db.AutoMigrate(&Metric{}, &Datapoint{}); err != nil {
		panic(err.Error())
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	return db
}
