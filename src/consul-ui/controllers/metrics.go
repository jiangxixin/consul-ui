package controllers

import (
	"consul-ui/models"
	"github.com/astaxie/beego"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"strconv"
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
		log.Fatal("consul client error : ", err)
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
		log.Fatal("consul client error : ", err)
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
	serviceName := c.GetString("service_name")
	if serviceName == "" {
		checked = false
		result.Code = 401
	}
	host := c.GetString("host")
	if host == "" {
		checked = false
		result.Code = 401
	}
	port := c.GetString("port")
	if port == "" {
		checked = false
		result.Code = 401
	}
	tag := c.GetString("tag")
	if tag == "" {
		checked = false
		result.Code = 401
	}
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
		log.Fatal("consul client error : ", err)
	}
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = serviceName
	registration.Name = serviceName
	registration.Port = portNum
	registration.Tags = []string{tag}
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
