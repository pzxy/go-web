package controllers

import (
	"github.com/astaxie/beego"
)

type DefaultController struct {
	beego.Controller
}

func (c *DefaultController) Get() {
	//通过配置文件的形式
	/*	c.Ctx.WriteString("appname:"+beego.AppConfig.String("appname")+
		"\nhttpport: "+beego.AppConfig.String("httpport")+
		"\nrunmode: "+beego.AppConfig.String("runmode"))*/

	//通过系统默认参数的形式
	//	c.Ctx.WriteString("\n\nappname:"+beego.AppName)

	//设置日志级别为info
	/*	beego.Trace("trace log test")
		beego.Info("info log test")
		beego.SetLevel(beego.LevelInformational)
		beego.Trace("trace log test2")*/

	//将数据输出到前端
	/*	c.TplName = "index.tpl"
		c.Data["Wrbsite"] = "beego.me"
		c.Data["Email"] = "wonkung@163.com"
		c.Data["TrueCond"] = true
		c.Data["FalseCOnd"] = false*/

	//结构体输出到前端
	/*c.TplName = "index.tpl"
	type u struct {
		Name string
		Age int
		Sex string
	}
	user := &u{
		Name:"wonkung",
		Age:24,
		Sex:"man",
	}
	c.Data["User"] = user*/
	//数组输出
	/*nums := []int{1,2,3,4,5,6,7,8,9,0}
	c.Data["Nums"] = nums*/
	//瞎定义输出 模版变量
	/*c.Data["Tplvar"] = "hey Diablo"*/
	//输出前端模版函数
	/*c.Data["Html"] = "<div>hello beego</div>"*/
	//pipe使用
	/*c.Data["Pipe"] = "<div>hello beego</div>"*/

	c.TplName = "home.html"
}
