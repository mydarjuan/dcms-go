package main

import (
	"github.com/astaxie/beego"
	"quickstart/models"
	_ "quickstart/routers"
)

func init() {
	models.RegisterDB()
}

func main() {
	beego.Run()
}
