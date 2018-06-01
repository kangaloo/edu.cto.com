package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalln(err)
		}

		go Conn(conn)
	}
}

// todo 结构体为参数时，什么情况传指针，什么情况传值
func Conn(c net.Conn) {
	buf := make([]byte, 1024)
	n, err := c.Read(buf)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(buf[:n])
}
