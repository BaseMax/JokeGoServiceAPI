package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Test code for checking docker-compose
	db, err := sql.Open("mysql", "user:pass@tcp(db:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("== Start Server ==")
	http.ListenAndServe(":8000", nil)
}
