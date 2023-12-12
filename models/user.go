package models

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/BaseMax/JokeGoServiceAPI/db"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique; not null" json:"username"`
	Password string `gorm:"not null" json:"password,omitempty"`
	Jokes    []Joke `gorm:"foreignKey:AuthorID"`
}

func hashPassword(password *string) {
	hash := sha512.Sum512([]byte(*password))
	*password = hex.EncodeToString(hash[:])
}

func RegisterUser(u *User) error {
	hashPassword(&u.Password)
	db := db.GetDB()
	return db.Create(&u).Error
}

func LoginUser(u *User) error {
	hashPassword(&u.Password)
	db := db.GetDB()
	return db.Where(u).First(&u).Error
}

func DeleteUserByName(name string) error {
	return db.GetDB().Delete(&User{}, "username = ?", name).Error
}
