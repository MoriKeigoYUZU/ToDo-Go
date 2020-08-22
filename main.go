package main

import (
	"log"
	"net/http"

	"github.com/ToDo-Go/handler"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/todos", handler.HomePage)
	router.GET("/pin", handler.Pin)
	//http.HandleFunc("/pin", handler.HomePage)
	log.Fatal(http.ListenAndServe(":8081", router))
}
