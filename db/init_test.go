package db

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	godotenv.Load("../.env")

	host := os.Getenv("DB_HOSTNAME")
	os.Setenv("DB_HOSTNAME", "invalidhost")
	assert.Error(t, Init())
	os.Setenv("DB_HOSTNAME", host)

	assert.NoError(t, Init())
	assert.NoError(t, Init())
}

func TestGetRandFunction(t *testing.T) {
	dbms := os.Getenv("DBMS")

	os.Setenv("DBMS", "mysql")
	assert.Equal(t, "RAND()", GetRandFunction())

	os.Setenv("DBMS", "postgres")
	assert.Equal(t, "RANDOM()", GetRandFunction())

	os.Setenv("DBMS", "invaliddb")
	assert.Empty(t, GetRandFunction())

	os.Setenv("DBMS", dbms)
}

func TestTruncateTable(t *testing.T) {
	type Model struct {
		ID   uint `gorm:"primaryKey"`
		Name string
	}

	dbms := os.Getenv("DBMS")

	db.AutoMigrate(&Model{})
	Init()
	db.Create(&Model{Name: "Name"})
	TruncateTable("models")
	model := &Model{Name: "Name"}
	db.Create(model)
	assert.NotEqual(t, 1, model.ID)

	os.Setenv("DBMS", dbms)
}

func TestGetDB(t *testing.T) {
	assert.NotEmpty(t, GetDB())
}
