package models

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/BaseMax/JokeGoServiceAPI/db"
)

const (
	FAKE_USER = "user"
	FAKE_PASS = "pass"
)

var d *gorm.DB

func TestMain(m *testing.M) {
	godotenv.Load("../.env")
	db.Init()
	d = db.GetDB()

	d.AutoMigrate(&User{}, &Joke{}, &Comment{})
	RegisterUser(&User{Username: FAKE_USER, Password: FAKE_PASS})

	code := m.Run()

	db.TruncateTable("comments", "jokes", "users")
	os.Exit(code)
}

func TestRegisterUser(t *testing.T) {
	u := User{Username: "newuser", Password: FAKE_PASS}
	assert.NoError(t, RegisterUser(&u))
	assert.Error(t, RegisterUser(&u))

}

func TestLoginUser(t *testing.T) {
	u := User{Username: "newuser", Password: FAKE_PASS}
	assert.NoError(t, LoginUser(&u))

	u = User{Username: "newuser", Password: "wrong"}
	assert.Error(t, LoginUser(&u))

	d.Delete(&u)
}

func TestDeleteUserByName(t *testing.T) {
	assert.NoError(t, DeleteUserByName("newuser"))
}
