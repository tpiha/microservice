package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	var db *gorm.DB
	var err error

	if db, err = gorm.Open(
		mysql.Open("root:root@tcp(database:3306)/microservice?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{}); err != nil {
		log.Printf("[initDb] error: %s", err)
	}

	if err := db.AutoMigrate(&Metric{}, &Datapoint{}); err != nil {
		panic(err.Error())
	}

	return db
}
