package database

import (
	"log"
	"os"
	"strings"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var connection *gorm.DB

func Connection() *gorm.DB {
	once.Do(func() {
		connection = initialize()
	})

	return connection
}

func initialize() *gorm.DB {
	driver := os.Getenv("DB_DRIVER")
	if strings.ToLower(driver) == "mysql" {
		return initializeMysql()
	}

	log.Fatal("Driver DB not specified")
	return nil
}
