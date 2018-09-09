package routers

import (
	"github.com/astaxie/beego"
	"go-study-beego/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//beego.Router("/customer/add", &controllers.CustomerController{})
	beego.Include(&controllers.MainController{})
	beego.AutoRouter(&controllers.MainController{})
	beego.AutoRouter(&controllers.CustomerController{})
}
