package routers

import (
	"github.com/astaxie/beego"
	"github.com/semihtok/KafkaBoard/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/kafka", &controllers.KafkaController{}, "post:MessageFromTopic")
	beego.Router("/actions", &controllers.ActionsController{})
}
