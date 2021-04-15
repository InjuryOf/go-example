package tcp_server

import (
	"fmt"
	"net"
	"sync"
)

var address = "localhost:9000"

func StartTcpServer() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		panic(fmt.Sprintf("net.Listen Error：%v\n", err))
	}
	defer listen.Close()
	fmt.Printf("Listen info：%v\n", address)
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("listen.Accept Error：%v\n", err)
		}
		fmt.Printf("请求信息：%s-%s\n", conn.RemoteAddr(), conn.LocalAddr())
		go logic(conn)
	}
}

var dbData string

// 逻辑处理
func logic(conn net.Conn) {
	defer conn.Close()

	// 业务逻辑处理
	message := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go handleServerWrite(conn, &wg, message)
	go handleServerRead(conn, &wg, message)
	wg.Wait()

	fmt.Println("请求处理完成")
}

func handleServerWrite(conn net.Conn, wg *sync.WaitGroup, message chan string) {
	defer wg.Done()

	// 读取请求数据
	buff := make([]byte, 1024)
	_, err := conn.Read(buff)
	if err != nil {
		fmt.Printf("conn.Read Error：%v\n", err)
	}
	dbData = string(buff)

	fmt.Println("写入业务数据")
	message <- "Write Done"
}

func handleServerRead(conn net.Conn, wg *sync.WaitGroup, message chan string) {
	defer wg.Done()
	select {
	case <-message:
		fmt.Printf("处理读请求：%s\n", dbData)
		conn.Write([]byte("success"))
	}
}
