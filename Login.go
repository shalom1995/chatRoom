package chat

import (
	"net"
	"time"
)

type Client struct {
	Addr string
	Name string
	C    chan string
}

var Onlin_Clients = make(map[string]Client)

var Message = make(chan string)

//	制造消息的方法
func MakeMessage(client Client, msg string) string {
	return "[" + client.Addr + "]|" + client.Name + ": " + msg
}

//	实现用户登录提醒的方法
func Login(conn net.Conn) {
	defer conn.Close()

	//	获取客户端地址
	clientAddr := conn.RemoteAddr().String()

	//	初始化客户
	client := Client{clientAddr, clientAddr, make(chan string)}

	//	新上线的客户加入map中
	Onlin_Clients[clientAddr] = client

	//	把消息写入用于广播消息的通道中
	Message <- MakeMessage(client, "login")

	go SendMsgToClient(conn, client)

	//for{
	//	;
	//}

	isQuit := make(chan bool)
	hasData := make(chan bool)

	go HandleMsgFromClient(conn, client, isQuit, hasData)

	for {
		select {
		case <-isQuit:
			close(client.C)
			delete(Onlin_Clients, clientAddr)
			Message <- MakeMessage(client, "log out")
			return
		case <-hasData:

		case <-time.After(45 * time.Second):
			delete(Onlin_Clients, clientAddr)
			Message <- MakeMessage(client, "time out leave")
			return
		}
	}
}
