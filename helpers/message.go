package main

import (
	handler "github.com/margostino/lumos/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Reply)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
