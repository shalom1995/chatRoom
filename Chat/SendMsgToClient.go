package chat

import (
	"net"
	"log"
	"fmt"
)

//	把 客户通道 的数据提取出来发给客户端
func SendMsgToClient(conn net.Conn, client Client) {
	for msg := range client.C {
		fmt.Println("client.C 接收到的数据：", msg)
		_, err := conn.Write([]byte(msg + "\n"))
		if err != nil {
			log.Panic(err)
		}
		//	为什么这里的conn是关闭的？因为conn所在的go程结束了
	}
}
