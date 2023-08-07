package db

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	godotenv.Load("../.env")
	assert.NoError(t, Init())
}

func TestGetDB(t *testing.T) {
	assert.NotEmpty(t, GetDB())
}
