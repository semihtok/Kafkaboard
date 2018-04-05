package controllers

import (
	"github.com/astaxie/beego"
	"github.com/semihtok/KafkaBoard/helpers"
)

type ChartController struct {
	beego.Controller
}

func (c *ChartController) Get() {
	c.Controller = helpers.LoadDefaultLayout(c.Controller, "chart")
}
