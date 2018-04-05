package main

import (
	"github.com/astaxie/beego"
	"github.com/semihtok/KafkaBoard/helpers"
	_ "github.com/semihtok/KafkaBoard/routers"
)

func main() {
	helpers.SetDefaultTemplateDelimeter("<<", ">>")
	beego.Run()
}
