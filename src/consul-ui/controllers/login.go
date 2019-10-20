package controllers

import (
	"consul-ui/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	result := new(models.JSONResult)
	checked := true
	username := c.GetString("username")
	if username == "" {
		checked = false
		result.Code = 401
	}
	password := c.GetString("password")
	if password == "" {
		checked = false
		result.Code = 401
	}
	if !checked {
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	if username == beego.AppConfig.String("auth_name") && password == beego.AppConfig.String("auth_passwd") {
		result.Code = 0
	} else {
		result.Code = 401
	}
	c.Data["json"] = result
	c.ServeJSON()
	return
}
