package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

//分类
type Category struct {
	Id              int64
	Title           string
	Created         time.Time
	Views           int64
	TopicTime       time.Time
	TopicCount      int64
	TopicLastUserId int64
}

//文章
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:size(1024)`
	Attachment      string
	Created         time.Time
	Updated         time.Time
	Views           int64
	Author          string
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
}

func init() {
	//register databases
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8")
	//2.注册表
	orm.RegisterModel(new(Category), new(Topic))
	//3.生成表 //force ture的话删了重建
	orm.RunSyncdb("default", false, true)
	//filepath.Ext()//获取文件的扩展名
}
func AddTopic(title, content string) error {
	o := orm.NewOrm()
	topic := &Topic{
		Title:   title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}
	_, err := o.Insert(topic)
	return err
}
func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := Category{Title: name}
	qs := o.QueryTable("category")
	qs = qs.Filter("title", name) //查询 category 表中的 title字段等于name的值

	if qs != nil {
		err := qs.One(cate)
		if err == nil { //没有报错 说明查到了 有重复的
			return err
		}
	}
	_, err := o.Insert(&cate)
	if err != nil {
		return err
	}
	return nil
}
func GetCategory() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 24)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	if err != nil {
		return err
	}
	return nil
}
