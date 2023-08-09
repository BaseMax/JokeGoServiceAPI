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

func TestGetDB(t *testing.T) {
	assert.NotEmpty(t, GetDB())
}
