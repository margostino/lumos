package main

import (
	handler "github.com/margostino/lumos/api"
	"github.com/margostino/lumos/db"
	"log"
	"net/http"
	"time"
)

func main() {
	variable := &db.Variable{
		Datetime:    time.Now().UTC().String(),
		Name:        "fire",
		Value:       "1",
		Observation: "nothing",
		Longitude:   10,
		Latitude:    15,
	}
	db.Append(variable)
	http.HandleFunc("/", handler.Reply)
	log.Println("Starting Lumos Server in :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
