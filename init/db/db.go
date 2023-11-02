package db

import (
	"fmt"
	"gin-dubbogo-consumer/global"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
	"time"
)

var DbPath = "./conf"
var DbFile = "k8s-client.db"

func Init() {
	if _, err := os.Stat(DbPath); err != nil {
		if err := os.MkdirAll(DbPath, os.ModePerm); err != nil {
			panic(fmt.Errorf("init db dir falied, err: %v", err))
		}
	}
	fullPath := DbPath + "/" + DbFile
	if _, err := os.Stat(fullPath); err != nil {
		if _, err := os.Create(fullPath); err != nil {
			panic(fmt.Errorf("init db file falied, err: %v", err))
		}
	}

	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	_ = db.Exec("PRAGMA journal_mode = WAL;")
	sqlDB, dbError := db.DB()
	if dbError != nil {
		panic(err)
	}
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	global.DB = db
	//global.LOG.Info("init db successfully")

}
