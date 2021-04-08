package main

import (
	"fmt"
	"net/http"
)

func main() {
	//单独写回调函数 http://127.0.0.1:8000/go
	http.HandleFunc("/go", myHandler)
	// addr:监听地址
	// handler：回调函数
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		fmt.Println("server start failed err:", err)
		return
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	//回复
	w.Write([]byte("www.baidu.com"))
}
