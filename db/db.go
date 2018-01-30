package db

import (
	"github.com/jrgirvan/devx/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB
var err error

// Init creates a connection to mysql database and
// migrates any new models
func Init(driver string, config map[string]string) {

	_, ok := config["ConnectionString"]
	if !ok {
		config["ConnectionString"] = "gorm.db"
	}

	db, err = gorm.Open(driver, config["ConnectionString"])
	if err != nil {
		log.Debug(err)
	}

	db.AutoMigrate(&models.Talk{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

//CloseDB ...
func CloseDB() {
	db.Close()
}
