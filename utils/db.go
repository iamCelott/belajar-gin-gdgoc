package utils

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB
var dbOnce sync.Once

func DB() *gorm.DB {
	dbOnce.Do(func() {
		dbHost := os.Getenv("DB_HOST")
		if dbHost == "" {
			panic("DB HOST not speciefy")
		}
		dbPort := os.Getenv("DB_PORT")
		if dbPort == "" {
			panic("DB PORT not speciefy")
		}
		dbName := os.Getenv("DB_NAME")
		if dbName == "" {
			panic("DB NAME not speciefy")
		}
		dbUser := os.Getenv("DB_USER")
		if dbUser == "" {
			panic("DB USER not speciefy")
		}
		dbPass := os.Getenv("DB_PASS")
		if dbPass == "" {
			panic("DB PASS not speciefy")
		}

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
		var err error
		dbInstance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("DB ERROR: %v", err))
		}
	})
	return dbInstance
}
