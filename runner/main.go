package main

import (
	handler "github.com/margostino/lumos/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Reply)
	log.Println("Starting Lumos Server in :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
