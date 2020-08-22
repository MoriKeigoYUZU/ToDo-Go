package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

//メモ
//w -> 書き出すため
//r -> ユーザからのリクエスト全て
//GOはバイナリでbodyが送られるのでJSONにエンコードする必要がある。

//受け取ったときの箱
type ReqBody struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Todos []Todo `json:"todos"`
}

type Todo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()
	buf := new(bytes.Buffer)                      //受け取り口
	if _, err := io.Copy(buf, body); err != nil { //bufにbody(r.body)をコピー
		return
	}
	var request ReqBody
	if err := json.Unmarshal(buf.Bytes(), &request); err != nil { //buf.Bytes() ->(マッピング) 構造体
		return
	}

	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Printf("%+v", request)
	fmt.Println("Endpoint Hit: homePage")

}

func main() {
	http.HandleFunc("/pin", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
