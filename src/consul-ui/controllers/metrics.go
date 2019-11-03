package controllers

import (
	"consul-ui/models"
	"github.com/astaxie/beego"
	consulapi "github.com/hashicorp/consul/api"
	"strconv"
	"strings"
)

type MetricsController struct {
	beego.Controller
}

func (c *MetricsController) Delete() {
	result := new(models.JSONResult)
	checked := true
	id := c.GetString("id")
	if id == "" {
		checked = false
		result.Code = 401
	}
	if !checked {
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	config := consulapi.DefaultConfig()
	config.Address = beego.AppConfig.String("consul_host")
	client, err := consulapi.NewClient(config)
	if err != nil {
		result.Code = 401
		result.Msg = err.Error()
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	err = client.Agent().ServiceDeregister(id)
	if err != nil {
		result.Code = 401
		result.Msg = err.Error()
	} else {
		result.Code = 0
		result.Msg = "success"
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *MetricsController) Get() {
	result := new(models.JSONResult)
	config := consulapi.DefaultConfig()
	config.Address = beego.AppConfig.String("consul_host")
	client, err := consulapi.NewClient(config)
	if err != nil {
		result.Code = 401
		result.Msg = err.Error()
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	services, err := client.Agent().Services()
	result.Code = 0
	result.Data = services
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *MetricsController) Post() {
	result := new(models.JSONResult)
	checked := true
	serviceName := strings.Trim(c.GetString("Service"), " ")
	if serviceName == "" {
		checked = false
		result.Code = 401
	}
	host := strings.Trim(c.GetString("Address"), " ")
	if host == "" {
		checked = false
		result.Code = 401
	}
	port := strings.Trim(c.GetString("Port"), " ")
	if port == "" {
		checked = false
		result.Code = 401
	}
	path := strings.Trim(c.GetString("Path"), " ")
	if path == "" {
		checked = false
		result.Code = 401
	}
	enable := strings.Trim(c.GetString("Enable"), " ")
	if enable == "" || (enable != "on" && enable != "Enable") {
		enable = "off"
	} else {
		enable = "on"
	}
	var meta map[string]string
	meta = make(map[string]string)
	meta["Path"] = path
	meta["Enable"] = enable
	tagStr := strings.Trim(c.GetString("Tags"), " ")
	if tagStr == "" {
		checked = false
		result.Code = 401
	}
	tags := strings.Split(tagStr, ",")
	portNum, err := strconv.Atoi(port)
	if err != nil {
		checked = false
		result.Code = 401
	}
	if !checked {
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	config := consulapi.DefaultConfig()
	config.Address = beego.AppConfig.String("consul_host")
	client, err := consulapi.NewClient(config)
	if err != nil {
		result.Code = 401
		result.Msg = err.Error()
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = serviceName
	registration.Name = serviceName
	registration.Port = portNum
	registration.Tags = tags
	registration.Meta = meta
	registration.Address = host
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		result.Code = 401
		result.Msg = err.Error()
	} else {
		result.Code = 0
		result.Msg = "success"
	}
	c.Data["json"] = result
	c.ServeJSON()
}
