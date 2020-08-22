package main

import (
	"fmt"
	"log"
	"net/http"
)

//w -> 書き出すため
//r -> ユーザからのリクエスト全て

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Printf("%+v", r)
	fmt.Println("Endpoint Hit: homePage")
}

func main() {
	http.HandleFunc("/pin", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
