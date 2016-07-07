package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"sort"
	"crypto/sha1"
	"net/http"
)

const(
	TOKEN = "weixin"
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

type WeChatController struct{
	beego.Controller
}

var w http.ResponseWriter
// type httpres struct{
// 	http.ResponseWriter
// }

func (c *WeChatController) Get(){
	//w :=new(httpres)
	var signature string = c.GetString("signature")
	beego.Info(signature)
	var timestamp string = c.GetString("timestamp")
	beego.Info(timestamp)
	var nonce string = c.GetString("nonce")
	beego.Info(nonce)
	var echostr string = c.GetString("echostr")
	beego.Info(echostr)
	beego.Info(sha1str(timestamp,nonce))
	if sha1str(timestamp,nonce)==signature{
		c.Ctx.WriteString(echostr)
	}else{
		c.Ctx.WriteString("")
	}
	c.TplName = "index.tpl"	
}

func sha1str(timestamp, nonce string) string {
	strs := sort.StringSlice{TOKEN, timestamp, nonce}
	sort.Strings(strs)
	str := ""
	for _, s := range strs {
		str += s
	}
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
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
