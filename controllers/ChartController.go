package controllers

import (
	"github.com/astaxie/beego"
	"github.com/semihtok/KafkaBoard/helpers"
)

// ChartController is chart controller as beego.Controller type
type ChartController struct {
	beego.Controller
}

// Get refers to "/chart" path page request (GET: /chart)
func (c *ChartController) Get() {
	c.Controller = helpers.LoadDefaultLayout(c.Controller, "chart")
}
