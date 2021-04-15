package tcp_server

import (
	"fmt"
	"net"
	"strconv"
)

func SendTcpRequest() {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(fmt.Sprintf("net.Dial Error：%v\n", err))
	}
	defer conn.Close()
	fmt.Printf("Connection info：%v\n", address)

	data := make(chan string)
	go handleClientWrite(conn, data)
	go handleClientRead(conn, data)

	fmt.Println(<-data)
	fmt.Println(<-data)

}

func handleClientWrite(conn net.Conn, data chan string) {
	for i := 0; i < 10; i++ {
		_, err := conn.Write([]byte(fmt.Sprintf("发起了一个TCP请求：%s", strconv.Itoa(i))))
		if err != nil {
			fmt.Printf("conn.Write err %v\n", err)
		}

	}
	data <- "请求发送结束"
}

func handleClientRead(conn net.Conn, data chan string) {
	buff := make([]byte, 1024)
	rlen, err := conn.Read(buff)
	if err != nil {
		fmt.Printf("conn.Read err %v\n", err)
	}
	fmt.Printf("返回数据：%v\n", string(buff[:rlen]))
	data <- "响应处理结束"
}
