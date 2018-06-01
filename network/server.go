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
func Conn(c net.Conn) ([]byte, error) {
	buf := make([]byte, 1024)
	_, err := c.Read(buf)
	if err != nil {
		return nil, err
	}
	fmt.Println(buf)
	return buf, nil
}
