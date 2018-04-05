package controllers

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
	helpers "github.com/semihtok/KafkaBoard/helpers"
	models "github.com/semihtok/KafkaBoard/models"
)

type KafkaController struct {
	beego.Controller
}

var kafkaResult []string

func (c *KafkaController) MessageFromTopic() {

	var topicRequestModel models.TopicRequestModel
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &topicRequestModel)

	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.ServeJSON()
	} else {
		go Read(topicRequestModel.Name)
		if kafkaResult == nil {
			log.Println(err)
			c.Ctx.Output.SetStatus(404)
		} else {
			response := models.TopicResponseModel{}
			response.Result = kafkaResult
			c.Data["json"] = response
			c.Ctx.Output.SetStatus(200)
			c.ServeJSON()
		}
	}
}

func Read(topicName string) {
	kafkaResult = helpers.ReadFromKafka(topicName)
}
