package systeminit

import "seasonCats/models"
//init函数用于main初始化使用
func init()  {
	//初始化数据库
	models.InitDb()
}
