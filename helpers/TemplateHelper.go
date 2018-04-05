package helpers

import (
	"github.com/astaxie/beego"
)

// SetDefaultTemplateDelimeter is changing default template delimeters with user defined delimeters
func SetDefaultTemplateDelimeter(startWith string, endWith string) {
	beego.BConfig.WebConfig.TemplateLeft = startWith
	beego.BConfig.WebConfig.TemplateRight = endWith
}

// LoadDefaultLayout is matching -> view & main layout page
func LoadDefaultLayout(c beego.Controller, templateName string) beego.Controller {
	c.Layout = "layouts/layout.tpl"
	c.TplName = templateName + ".tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Styles"] = "layouts/styles.tpl"
	c.LayoutSections["Scripts"] = "layouts/scripts.tpl"
	return c
}
