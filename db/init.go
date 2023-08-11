package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Init() error {
	var err error
	var dialector gorm.Dialector

	if db != nil {
		return nil
	}

	conf := &gorm.Config{}
	if os.Getenv("DEBUG") != "true" {
		conf = &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}
	}

	switch os.Getenv("DBMS") {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("DB_HOSTNAME"), os.Getenv("MYSQL_DATABASE"),
		)
		dialector = mysql.Open(dsn)
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			os.Getenv("DB_HOSTNAME"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"),
		)
		dialector = postgres.Open(dsn)
	}

	db, err = gorm.Open(dialector, conf)
	if err != nil {
		db = nil
		return err
	}

	return nil
}

func GetRandFunction() string {
	switch os.Getenv("DBMS") {
	case "mysql":
		return "RAND()"
	case "postgres":
		return "RANDOM()"
	}
	return ""
}

func GetDB() *gorm.DB {
	return db
}
