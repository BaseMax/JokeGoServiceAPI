package main

import (
	"log"
	"os"

	"github.com/BaseMax/JokeGoServiceAPI/db"
	"github.com/BaseMax/JokeGoServiceAPI/migration"
	"github.com/BaseMax/JokeGoServiceAPI/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	if err := migration.Init(); err != nil {
		log.Fatal(err)
	}
	r := routes.Init()

	addr := os.Getenv("RUNNING_ADDR")
	if addr == "" {
		addr = ":8000"
	}
	r.Logger.Fatal(r.Start(addr))
}
