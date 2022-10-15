package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/http2"
)

func main() {
	server := http.Server{}
	http2.ConfigureServer(&server, &http2.Server{})
	http.HandleFunc("/post", post)
	http.HandleFunc("/getatob", getatob)
	http.HandleFunc("/getbtoa", getbtoa)
	http.ListenAndServeTLS(":8081", "", "", nil)
}

var atob = ""
var btoa = ""

func post(w http.ResponseWriter, request *http.Request) {
	//POSTを取得
	request.ParseForm()
	if len(request.Form) == 0 {
		fmt.Println("Empty Request")
		return
	}
	if request.Method != "POST" {
		return
	}
	println(request.Method)
	//リクエストからプログラムを取得し保存
	text := strings.Join(request.Form["text"][:], "")
	user := strings.Join(request.Form["user"][:], "")
	timestamp := strings.Join(request.Form["time"][:], "")
	println(timestamp + ":user:" + user + "text:" + text)
	if user == "a" {
		atob = text
	} else {
		btoa = text
	}
	w.Write([]byte("hello"))
	//ファイルを閉じる
}

func getatob(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var result = atob
	w.Write([]byte(result))
}

func getbtoa(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var result = btoa
	w.Write([]byte(result))
}
