package models

import (
	"os"
	"testing"

	"github.com/BaseMax/JokeGoServiceAPI/db"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var d *gorm.DB

func TestMain(m *testing.M) {
	godotenv.Load("../.env")
	db.Init()
	d = db.GetDB()
	d.AutoMigrate(&User{})
	code := m.Run()
	os.Exit(code)
}

func TestRegisterUser(t *testing.T) {
	u := User{Username: "user", Password: "pass"}
	assert.NoError(t, RegisterUser(&u))
	assert.Error(t, RegisterUser(&u))

}

func TestLoginUser(t *testing.T) {
	u := User{Username: "user", Password: "pass"}
	assert.NoError(t, LoginUser(&u))

	u = User{Username: "user", Password: "wrong"}
	assert.Error(t, LoginUser(&u))

	d.Delete(&u)
}

func TestDeleteUserByName(t *testing.T) {
	assert.NoError(t, DeleteUserByName("user"))
}
