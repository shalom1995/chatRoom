package chat

import (
	"net"
	"fmt"
	"log"
)

func HandleMsgFromClient(conn net.Conn, client Client,isQuit,hasData chan bool) {


	buf := make([]byte, 4098)

	for {
		//	读取客户端发来的数据，通过读取数据的长度来判断客户端是否在线
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Printf("客户%s下线\n", client.Name)
			isQuit <- true
		}
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("====================")
		//fmt.Println(string(buf))
		//	将读取的数据转成字符串，但需要先去除换行符 rename|
		msg := string(buf[:n-1])
		fmt.Println(msg)

		if len(msg) == 3 && msg == "who" {
			conn.Write([]byte("client list:\n"))

			for _, cli := range Onlin_Clients {
				msg := cli.Addr + ":" + cli.Name + "\n"
				conn.Write([]byte(msg))
			}
		} else if len(msg) > 7 && msg[:7] == "rename|" {
			conn.Write([]byte("rename success!\n"))
			client.Name = msg[7:]
			//client.Name = strings.Split(msg,"|")[1]

			Onlin_Clients[client.Addr] = client
		}else if len(msg) == 4 && msg == "exit"{
			isQuit<-true
		} else {
			//conn.Write([]byte(msg))
			Message<-MakeMessage(client,msg)
		}
		hasData<-true
	}
}

//delete(Onlin_Clients,client.Addr)
