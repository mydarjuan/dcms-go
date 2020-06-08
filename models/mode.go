package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type Student struct {
	Id   int64
	Name string
	No   string `orm:"index"`
}

func RegisterDB() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(Student))
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RunSyncdb("default", false, true)
}
