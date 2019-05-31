package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(w http.ResponseWriter, r *http.Request)

func main() {
	//1自己构造server
	server := http.Server{
		Addr:              ":8080",
		Handler:           &myHandler3{}, //自己实现server  根据路由调用方法
		ReadHeaderTimeout: 5 * time.Second,
	}
	//创建空间
	mux = make(map[string]func(w http.ResponseWriter, r *http.Request))
	mux["/hello"] = sayHello3
	mux["/bey"] = sayBey3
	//2监听请求
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

//3监听到以后根据mux来获取请求的方法 这里自己实现server
type myHandler3 struct{}

func (*myHandler3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
}
func sayHello3(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "say hello")
}

func sayBey3(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "say bey")
}
