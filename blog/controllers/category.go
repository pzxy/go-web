package controllers

import (
	"github.com/astaxie/beego"
	"myapp/models"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	op := this.GetString("op")
	switch op {
	case "add":
		name := this.GetString("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	case "del":
		id := this.GetString("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	}
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsCategory"] = true
	this.TplName = "category.html"

	var err error
	this.Data["Category"], err = models.GetCategory()
	if err != nil {
		beego.Error(err)
	}
}
