package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ToDo-Go/db/mysql"

	"github.com/julienschmidt/httprouter"
)

//メモ
//w -> 書き出すため
//r -> ユーザからのリクエスト全て
//GOはバイナリでbodyが送られるのでJSONにエンコードする必要がある。

//受け取ったときの箱

type Todo struct {
	Title string `json:"title"`
}

func CreateTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body := r.Body
	defer body.Close()
	buf := new(bytes.Buffer)                      //受け取り口
	if _, err := io.Copy(buf, body); err != nil { //bufにbody(r.body)をコピー
		return
	}

	//UnmarshalJSONから構造体
	var request Todo
	if err := json.Unmarshal(buf.Bytes(), &request); err != nil { //buf.Bytes() ->(マッピング) 構造体
		return
	}

	_, err := mysql.DB.Exec("INSERT INTO todo(title) VALUES (?)", request.Title)
	if err != nil {
		log.Fatalln(err)
	}

	//w : ユーザに流したい情報をセットできる
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Encode 構造体からJSON
	//NewEncoder(w)勝手に流してくれる
	if err := json.NewEncoder(w).Encode(request); err != nil {
		panic(err)
	}
}

func Pin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello")
}
