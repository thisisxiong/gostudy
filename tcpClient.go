package main

import (
	"bufio"
	"fmt"
	"net"
	"simple/proto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("err:", err)
	}

	//defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := fmt.Sprintf("hello How are you %v", i)
		message, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("err：", err)
			return
		}
		conn.Write([]byte(message))

		reader := bufio.NewReader(conn)
		recv, err := proto.Decode(reader)
		if err != nil {
			fmt.Println("recv err:", err)
			return
		}
		fmt.Println("recv:", recv)

	}
	//inputReader := bufio.NewReader(os.Stdin)
	//for {
	//	input,_ := inputReader.ReadString('\n') //读取用户输入
	//	inputInfo := strings.Trim(input,"\r\n")
	//	if strings.ToUpper(inputInfo) == "Q"{
	//		return
	//	}
	//	_,err := conn.Write([]byte(inputInfo)) //发送数据
	//	if err != nil{
	//		return
	//	}
	//	buf := [512]byte{}
	//	n,err := conn.Read(buf[:])
	//	if err != nil{
	//		fmt.Println("recv err:",err)
	//		return
	//	}
	//	fmt.Println(string(buf[:n]))
	//}

}
