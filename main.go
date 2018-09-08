package main

import (
	"net"
	"log"
	"./chat"
	"fmt"
)

func main() {
	listener,err:=net.Listen("tcp","127.0.0.1:8002")
	if err!=nil{
		log.Panic(err)
	}

	go chat.Spread()

	for{
		conn,err:=listener.Accept()
		fmt.Println("客户端已连接...")
		if err!=nil{
			log.Panic(err)
		}

		go chat.Login(conn)

	}
}
