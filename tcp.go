package main

import (
	"bufio"
	"fmt"
	"net"
	"simple/proto"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		recvStr, err := proto.Decode(reader) //读取数据
		if err != nil {
			fmt.Println("接收消息 err:", err)
			break
		}
		fmt.Println("收到client端发来的数据：", recvStr)

		ret, err := proto.Encode(recvStr)
		if err != nil {
			fmt.Println("返回消息 err：", err)
			break
		}
		conn.Write(ret) //发送数据
	}

}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("listent failed,err:", err)
		return
	}
	fmt.Println("=== server is running listen: 127.0.0.1:8889 ===")
	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("accept failed,err", err)
			continue
		}
		go process(conn)
	}
}
