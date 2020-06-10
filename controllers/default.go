package controllers

import (
	"dcms-go/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	o := orm.NewOrm()
	var students []models.Student
	o.Raw("SELECT id,name,no FROM student").QueryRows(&students)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Students"] = students
	c.TplName = "index.tpl"
}

func (c *MainController) GetStudents() {
	o := orm.NewOrm()
	var students []models.Student
	o.Raw("SELECT id,name,no FROM student").QueryRows(&students)
	c.Data["json"] = students
	c.ServeJSON()
}
