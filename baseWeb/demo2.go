package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	//mux控制 http的访问
	mux.HandleFunc("/hello", sayHello2)

	//获取当前相对路径
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//静态相对路径
	mux.Handle("/static/",
		http.StripPrefix("/static/", //将前面的static覆盖掉
			http.FileServer(http.Dir(wd)))) //定位wd中的路径

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL: "+r.URL.String())
}
func sayHello2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Let it go")
}
