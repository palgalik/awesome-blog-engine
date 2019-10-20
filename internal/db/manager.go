package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
	"path/filepath"
)

var instance *gorm.DB
var isInitialized = false

func GetInstance() *gorm.DB {
	if isInitialized {
		return instance
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	rootDir := filepath.Dir(wd)
	db, err := gorm.Open("sqlite3", rootDir+"/awesome-blog-engine.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	instance, isInitialized = db, true

	return instance
}
