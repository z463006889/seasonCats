package models

import (
	"seasonCats/utils"
	"github.com/astaxie/beego/orm"
)
//模型注册
func init()  {
	orm.RegisterModel(new(Users))
}

type Users struct {
	Id int
	LoginName string `orm:"size(64)"`
	Password string	`orm:"size(40)"`
	NickName string	`orm:"size(64)"`
}
//创建用户
func CreateUser(loginName string,pwd string,nickName string)(int64,error)  {
	 o:=orm.NewOrm()
	 user:=Users{
	 	LoginName:loginName,
	 	Password:utils.Md5(pwd),
	 	NickName:nickName,
	 }
	num,err:=o.Insert(&user)
	if err!=nil {
		return -1,err
	}
	return num,nil
}
//获取用户的基本信息
func GetUserBasic(userId int)  error{
	o:=orm.NewOrm()
	user:=Users{Id:userId}
	err :=o.Read(&user)
	if err!=nil {
		return err
	}
	return nil
}
//删除用户
func DeleteUser(userId int) error {
	o:=orm.NewOrm()
	user:=Users{Id:userId}
	if _,err:=o.Delete(&user);err!=nil{
		return err
	}
	return nil
}

//修改用户昵称
func UpdateNickName(userId int,nickName string) error  {
	o:=orm.NewOrm()
	user:=Users{Id:userId}
	if err:=o.Read(&user);err==nil{
		user.NickName=nickName
		_,err=o.Update(&user,"nickName")  //修改用户的昵称
		if err!=nil {
			return err
		}
	}
	return nil
}


