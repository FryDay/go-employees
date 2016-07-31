package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/FryDay/go-employees/models"
)

var db *sql.DB

func main() {
	newDB, err := models.NewDB("./emps.db")
	if err != nil {
		log.Panic(err)
	}
	db = newDB
	defer db.Close()

	router := newRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
