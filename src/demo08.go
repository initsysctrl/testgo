/**
 * @Author: yp
 * @Date: 2021/7/30 5:24 下午
 * @Des:go tcp server
**/
package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	for {
		reader := bufio.NewReader(conn)
		var buff [128]byte
		n, err := reader.Read(buff[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		message := string(buff[:n])
		fmt.Println("receive message", message)
		_, _ = conn.Write([]byte(message)) //send data

	}

}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println()
		return
	}
	go process(conn)

}
