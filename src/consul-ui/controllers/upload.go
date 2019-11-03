package controllers

import (
	"consul-ui/models"
	"github.com/astaxie/beego"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/tealeg/xlsx"
	"strconv"
	"strings"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Post() {
	result := new(models.JSONResult)
	f, h, _ := c.GetFile("file")
	path := beego.AppConfig.String("upload_tmp_path") + h.Filename
	c.SaveToFile("file", path)
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		result.Code = 401
		result.Msg = err.Error()
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	defer f.Close()
	sheet := xlFile.Sheets[0]
	for index, row := range sheet.Rows {
		if index == 0 {
			continue
		}
		serviceName := strings.Trim(row.Cells[0].String(), " ")
		host := strings.Trim(row.Cells[1].String(), " ")
		portNum, _ := strconv.Atoi(row.Cells[2].String())
		tags := strings.Split(strings.Trim(row.Cells[3].String(), " "), ",")
		var meta map[string]string
		meta = make(map[string]string)
		meta["Path"] = strings.Trim(row.Cells[4].String(), " ")
		enable := strings.Trim(row.Cells[5].String(), " ")
		if enable == "" || (enable != "on" && enable != "Enable") {
			enable = "off"
		} else {
			enable = "on"
		}
		meta["Enable"] = enable
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
	}
	result.Code = 0
	result.Msg = "success"
	c.Data["json"] = result
	c.ServeJSON()
}
