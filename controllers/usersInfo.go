package controllers

import (
	"fmt"
)

type UserControllers struct {
	baseController
}

type userInfo struct {
	userName string `json:"user_name"`
	pwd string `json:"password"`
}


func (u UserControllers)Prepare()  {
	
}

func (u *UserControllers) CreateUser() {
	username:=u.GetString("username")
	username=u.Ctx.Input.Param(":username")
	fmt.Print(username)
	u.Data["json"]= &userInfo{"zhaohu","123456"}
	u.ServeJSON()
}

func (u *UserControllers)login(username string,pwd string)  {

}
