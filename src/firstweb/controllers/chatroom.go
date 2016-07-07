package controllers

import (
	//"encoding/json"
    "github.com/gorilla/websocket"
    "fmt"
    "github.com/astaxie/beego"
    "net/http"
    //"firstweb/models"
    "container/list"
    "time"
)

type ChatRoomController struct{
	beego.Controller
}

type Event struct {
	User      string
	Timestamp int // Unix timestamp (secs)
	Content   string
}

type Subscriber struct {
	Conn *websocket.Conn
	Send chan Event // Only for WebSocket users; otherwise nil.
}

// var(
// 	subscribe *Subscriber= &Subscriber{}
// 	publish = make(chan models.Event, 10)
// 	subscribers = list.New()
// )
var users=list.New()

func newEvent(user, msg string) Event {
	return Event{ user, int(time.Now().Unix()), msg}
}

func (this *ChatRoomController) Say(){
	uname:= this.GetString("username")
	this.TplName = "chatroom.tpl"
	this.Data["Username"] = uname
	beego.Info(uname)
}

func (this *ChatRoomController) Welcome(){
	this.TplName = "welcome.tpl"
}

func (this *ChatRoomController) Chat(){
	//go Echo(&websocket.Conn)
	//http.Handle("/",websocket.Handler(Echo))

	uname:= this.GetString("username")
	beego.Info(uname)
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	users.PushBack(ws)
	subscribe:=&Subscriber{
		Conn : ws,
		Send : make(chan Event,10),
	}
	//item:=subscribers.PushFront(ws)
	//defer subscribers.Remove(item)
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println("Received back from client: " + string(p))
		subscribe.Send <- newEvent(uname, string(p))
		go subscribe.chatroom()
	}

	
}

// func init(){
// 	go chatroom()
// }

func (this *Subscriber)chatroom(){
	for b:= range this.Send{
		for item := users.Front(); item != nil; item = item.Next(){
			ws, ok := item.Value.(*websocket.Conn)
        	if !ok {
            	panic("item not *websocket.Conn")
        	}
        	// if item.Value == this.Conn {
         //    	continue
        	// }
        	err:=ws.WriteJSON(b)
			if err!=nil{
				break
			}
		}
		//data:= json.Marshal(b)
		
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
