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

	//UnmarshalJSONから構造体
	if err := json.Unmarshal(buf.Bytes(), &request); err != nil { //buf.Bytes() ->(マッピング) 構造体
		return
	}

	//w : ユーザに流したい情報をセットできる
	w.Header().Set("Content-Type", "application/json; charset=UTF-8") // <- Added
	w.Header().Set("Set-Cookie", "sessionId=38afes7a8")
	w.Header().Set("hoge", "sessionId=38afes7a8")
	w.WriteHeader(http.StatusOK) // <- Added

	//Encode 構造体からJSON
	//NewEncoder(w)勝手に流してくれる
	if err := json.NewEncoder(w).Encode(request); err != nil {
		panic(err)
	}

	fmt.Printf("%+v", request)
	fmt.Println("Endpoint Hit: homePage")

}

func main() {
	http.HandleFunc("/pin", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
