package main

import (
	"log"
	"net/http"

	"gopkg.in/gorp.v1"

	"github.com/FryDay/go-employees/models"
)

var db *gorp.DbMap

func main() {
	newDB, err := models.NewDB("./emps.db")
	if err != nil {
		log.Panic(err)
	}
	db = newDB
	defer db.Db.Close()

	router := newRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
