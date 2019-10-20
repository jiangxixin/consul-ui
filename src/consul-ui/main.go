package main

import (
	_ "consul-ui/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
