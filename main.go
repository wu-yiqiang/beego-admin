package main

import (
	_ "beego-admin/routers"
	_ "beego-admin/models"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}