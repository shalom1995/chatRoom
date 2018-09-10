package chat

func Spread() {
	for { //	for循环不断监听
		//fmt.Println("for循环不断监听 Message 里的数据...")
		msg := <-Message //	接收消息
		//fmt.Println("Message接收到消息...",msg)
		for _, cli := range Onlin_Clients {
			cli.C <- msg //	遍历所有在线用户，广播消息
		}
	}
}
