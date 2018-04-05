package controllers

import (
	"github.com/astaxie/beego"
	"github.com/semihtok/KafkaBoard/helpers"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Controller = helpers.LoadDefaultLayout(c.Controller, "index")
}
