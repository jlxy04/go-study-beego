package controllers

import (
	"github.com/astaxie/beego"
	"go-study-beego/models"
	"fmt"
)

type CustomerController struct {
	beego.Controller
}

func (this *CustomerController) Add() {
	c := models.Customer{}
	err := this.ParseForm(&c)
	if err != nil {
		beego.Error("解析到表单异常")
	}
	fmt.Println(&c)
	models.InsertCustomer(&c)
	this.TplName = "index.tpl"
}
