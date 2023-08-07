package db

import (
	"fmt"
	"os"

	"github.com/BaseMax/JokeGoServiceAPI/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("DB_HOSTNAME"), os.Getenv("MYSQL_DATABASE"),
	)
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}

	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		return err
	}

	err = db.AutoMigrate(&models.User{}, &models.Joke{}, &models.Comment{})
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
