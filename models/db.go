package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

func InitDb()  {
	dbType := beego.AppConfig.String("db_type")
	//连接名称
	dbAlias := beego.AppConfig.String("db_alias")
	//数据库名称
	dbName := beego.AppConfig.String("db_name")
	//数据库连接用户名
	dbUser := beego.AppConfig.String("db_user")
	//数据库连接用户名
	dbPwd := beego.AppConfig.String("db_pwd")
	//数据库IP（域名）
	dbHost := beego.AppConfig.String("db_host")
	//数据库端口
	dbPort := beego.AppConfig.String("db_port")
	//数据库编码格式
	dbCharset := beego.AppConfig.String("db_charset")

	dbSource:= dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset
	//orm.RegisterDriver(dbType,orm.DRMySQL)
	orm.RegisterDataBase(dbAlias,dbType,dbSource)
	orm.RunSyncdb(dbAlias,false,true)
}
