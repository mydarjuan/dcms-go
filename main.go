package main

import (
	"dcms-go/models"
	_ "dcms-go/routers"
	"github.com/astaxie/beego"
)

func init() {
	models.RegisterDB()
}

func main() {
	beego.Run()
}
