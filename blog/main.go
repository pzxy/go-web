package main

import (
	"github.com/astaxie/beego"
	_ "myapp/models"
	_ "myapp/routers"
)

func main() {
	beego.Run()
}
