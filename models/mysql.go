package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
const (
	_DB_NAME    = "default"
	_SQL_DRIVER = "mysql"
)

func init() {
	orm.RegisterModel(new(Customer))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(_DB_NAME, _SQL_DRIVER, "root:dzy_wx_123456@tcp(120.76.157.34:3306)/test?parseTime=true", 10)
}
