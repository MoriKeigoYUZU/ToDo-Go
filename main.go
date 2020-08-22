package main

import (
	"log"
	"net/http"

	"github.com/ToDo-Go/handler"
)

func main() {
	http.HandleFunc("/pin", handler.HomePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
