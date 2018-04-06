package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/semihtok/KafkaBoard/helpers"
	models "github.com/semihtok/KafkaBoard/models"
)

// ActionsController is actions controller as beego.Controller type
type ActionsController struct {
	beego.Controller
}

// Get refers to "/actions" path page request (GET: /actions)
func (c *ActionsController) Get() {
	c.Controller = helpers.LoadDefaultLayout(c.Controller, "actions")
}

// Post refers to "/actions" path post request (Post: /actions)
func (c *ActionsController) Post() {
	var topicRequestModel models.TopicRequestModel
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &topicRequestModel)

	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.ServeJSON()
	} else {
		status := helpers.WriteToKafka(topicRequestModel.Name, []byte(topicRequestModel.Message), 0)

		if status {
			c.Ctx.Output.SetStatus(200)
			c.ServeJSON()
		} else {
			c.Ctx.Output.SetStatus(500)
			c.ServeJSON()
		}

	}
}
