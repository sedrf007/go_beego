package routers

import (
	"firstweb/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/trysql",&controllers.SqlController{})
    beego.Get("/hello",func(ctx *context.Context){
    ctx.Output.Body([]byte("胖超你好"))
})
    beego.Router("/login",&controllers.LoginController{})
    beego.Router("/say",&controllers.ChatRoomController{},"*:Say")
    beego.Router("/chat",&controllers.ChatRoomController{},"*:Chat")
}


