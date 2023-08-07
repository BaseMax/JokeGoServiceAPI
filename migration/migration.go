package migration

import (
	"github.com/BaseMax/JokeGoServiceAPI/db"
	"github.com/BaseMax/JokeGoServiceAPI/models"
)

func Init() error {
	return db.GetDB().AutoMigrate(&models.User{}, &models.Joke{}, &models.Comment{})
}
