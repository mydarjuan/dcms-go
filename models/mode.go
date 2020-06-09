package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"strconv"
)

type Student struct {
	Id   int64
	Name string
	No   string `orm:"index"`
}

func RegisterDB() {
	orm.RegisterModel(new(Student))
	orm.Debug, _ = strconv.ParseBool(beego.AppConfig.String("debug"))
	_ = orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("database_url"), 30)
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RunSyncdb("default", false, true)
}
