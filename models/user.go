package models

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/BaseMax/JokeGoServiceAPI/db"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique; not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

func hashPassword(password *string) {
	hash := sha512.Sum512([]byte(*password))
	*password = hex.EncodeToString(hash[:])
}

func RegisterUser(u *User) error {
	hashPassword(&u.Password)
	db := db.GetDB()
	err := db.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func LoginUser(u *User) error {
	hashPassword(&u.Password)
	db := db.GetDB()
	if err := db.Where(u).First(&u).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUserByName(name string) error {
	return db.GetDB().Delete(&User{}, "username = ?", name).Error
}
