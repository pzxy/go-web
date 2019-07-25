package controllers

import (
	"github.com/astaxie/beego"
	"myapp/models"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	title := this.GetString("title")
	content := this.GetString("content")
	err := models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}
}
func (this *TopicController) Add() {
	this.TplName = "topic_add.html"
}
