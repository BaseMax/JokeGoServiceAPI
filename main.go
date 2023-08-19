package main

import (
	"log"
	"os"

	"github.com/BaseMax/JokeGoServiceAPI/db"
	"github.com/BaseMax/JokeGoServiceAPI/models"
	"github.com/BaseMax/JokeGoServiceAPI/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	err := db.GetDB().AutoMigrate(&models.User{}, &models.Joke{}, &models.Comment{})
	if err != nil {
		log.Fatal(err)
	}
	r := routes.Init()

	addr := os.Getenv("RUNNING_ADDR")
	if addr == "" {
		addr = ":8000"
	}
	r.Logger.Fatal(r.Start(addr))
}
