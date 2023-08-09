package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Init() error {
	if db != nil {
		return nil
	}

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("DB_HOSTNAME"), os.Getenv("MYSQL_DATABASE"),
	)
	conf := &gorm.Config{}
	if os.Getenv("DEBUG") != "true" {
		conf = &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}
	}
	db, err = gorm.Open(mysql.Open(dsn), conf)
	if err != nil {
		db = nil
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
