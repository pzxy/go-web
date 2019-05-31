package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	//设置路由
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

//第一个参数是接口，所以没有值传递还是地址传递的区别，第二个是结构体，实例化以后是有存储空间的，如果用值传递的话效率会低，所以要地址传递的
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Let it go") //这里ResponseWriter实现了Writer的中的write（）方法 等于实现了Write接口 所以能当做参数传进去
}
