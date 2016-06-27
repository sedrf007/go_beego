package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type LoginController struct{
	beego.Controller
}

type SqlController struct{
	beego.Controller
}

type user struct{
	Name string `form:"nickname"`
	password string 
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *LoginController) Get(){
	c.TplName = "login.tpl"
}

func (c *SqlController) Post(){
	c.Data["nickname"]=c.GetString("nickname")
	c.Data["password"]=c.GetString("password")
	
	//beego.Debug(fmt.Println(d))
	//c.Data["nickname"]=v.Type()
	c.TplName = "sql.tpl"
}
