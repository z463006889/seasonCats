package main

import (
	_ "seasonCats/routers"
	_"seasonCats/systeminit"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

