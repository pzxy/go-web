package routers

import (
	"github.com/astaxie/beego"
	"myapp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "get:Get;post:Post")
	beego.Router("/category", &controllers.CategoryController{}, "get:Get;post:Post")
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/topic", &controllers.TopicController{}, "get:Get;post:Post")
}
