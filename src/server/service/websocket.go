package service

import (
	"github.com/kataras/iris"
	"fmt"
)

type clientPage struct {
	Title string
	Host  string
}

func Websocket() {
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("main.html", clientPage{"Client Page", ctx.HostString()})
	})

	iris.Config.Websocket.Endpoint = "/doubeles"

	var myChatRoom = "room1"
	iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {
		//加入房间
		c.Join(myChatRoom)

		c.On("chat", func(message string) {
			fmt.Println(message)
			c.To(myChatRoom).Emit("chat", message)
		})

		c.OnDisconnect(func() {
			fmt.Println("disconnect...")
		})
	})

	iris.Listen(":8080")
}
