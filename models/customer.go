package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"go-study-beego/util"
)

type Customer struct {
	CustomerId string `form:"-" orm:"pk"`
	Name       string `form:"name"`
	Age        int `form:"age"`
	Sex        string `form:"sex"`
	Birthday   time.Time
	Mobile     string `form:"mobile"`
	Remark     string
	CreateTime time.Time
}

func InsertCustomer(c *Customer) *Customer {
	o := orm.NewOrm()
	c.CustomerId = common.GetUUID()
	fmt.Println(*c)
	id, err := o.Insert(c)
	fmt.Println(id, err)
	return c
}
