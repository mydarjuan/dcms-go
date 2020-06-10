package routers

import (
	"dcms-go/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index/students", &controllers.MainController{}, "get:GetStudents")
}
