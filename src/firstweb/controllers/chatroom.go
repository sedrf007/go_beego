package controllers

import (
    "github.com/gorilla/websocket"
    "fmt"
    "github.com/astaxie/beego"
    "net/http"
)

type ChatRoomController struct{
	beego.Controller
}

func (this *ChatRoomController) Say(){
	this.TplName = "chatroom.tpl"
}

func (this *ChatRoomController) Chat(){
	//go Echo(&websocket.Conn)
	//http.Handle("/",websocket.Handler(Echo))
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println("Received back from client: " + string(p))
	}
}


// func Echo(ws *websocket.Conn) {
//     var err error

//     for {
//         var reply string

//         if err = websocket.Message.Receive(ws, &reply); err != nil {
//             fmt.Println("Can't receive")
//             break
//         }

//         fmt.Println("Received back from client: " + reply)

//         msg := "Received:  " + reply
//         fmt.Println("Sending to client: " + msg)

//         if err = websocket.Message.Send(ws, msg); err != nil {
//             fmt.Println("Can't send")
//             break
//         }
//     }
// }
