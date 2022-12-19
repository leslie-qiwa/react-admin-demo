package database

import (
	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var (
	DB    *gorm.DB
	err   error
	DBErr error
)

// DBConnection create database connection
func DBConnection(dbName string) error {
	var db = DB

	logMode := viper.GetBool("DB_LOG_MODE")
	loglevel := logger.Silent
	if logMode {
		loglevel = logger.Info
	}

	sdb := sqlite.Open(dbName)
	db, err = gorm.Open(sdb, &gorm.Config{
		Logger: logger.Default.LogMode(loglevel),
	})

	if err != nil {
		DBErr = err
		log.Println("Db connection error")
		return err
	}

	err = db.AutoMigrate(migrationModels...)

	if err != nil {
		return err
	}
	DB = db

	return nil
}

// GetDB connection
func GetDB() *gorm.DB {
	return DB
}

// GetDBError connection error
func GetDBError() error {
	return DBErr
}
