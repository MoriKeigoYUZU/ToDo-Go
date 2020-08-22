package main

import (
	"log"
	"net/http"

	"github.com/ToDo-Go/db/mysql"
	"github.com/ToDo-Go/handler"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load .env file: %v", err)
	}

	db := mysql.ConnectLocalSQL()
	defer db.Close()

	router := httprouter.New()
	router.POST("/todos", handler.HomePage)
	router.GET("/pin", handler.Pin)
	log.Fatal(http.ListenAndServe(":8081", router))
}
