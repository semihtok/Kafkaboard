package controllers

import (
	"github.com/astaxie/beego"
	"github.com/semihtok/KafkaBoard/helpers"
)

// MainController is root page controller as beego.Controller type
type MainController struct {
	beego.Controller
}

// Get refers to root path for page request (GET: /chart)
func (c *MainController) Get() {
	c.Controller = helpers.LoadDefaultLayout(c.Controller, "index")
}
