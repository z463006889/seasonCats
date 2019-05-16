package routers

import (
	"seasonCats/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    //user
    beego.Router("users/addUser/:username",&controllers.UserControllers{},"get:CreateUser")
}
