package main

import (
	"fmt"
	"io"
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
		log.Printf("%s connected .", conn.RemoteAddr())

		go Conn(conn)
	}
}

// todo 结构体为参数时，什么情况传指针，什么情况传值
func Conn(c net.Conn) {
	defer c.Close()

	all := make([]byte, 0, 1024)
	buf := make([]byte, 2)

	for {

		// 读完最后的数据后，会阻塞
		n, err := c.Read(buf)

		if err != nil {

			// 从socket读取完所有数据的当次，或者下一次会得到 io.EOF
			// 目前没有出现在当次，又因为当次读取完成后，下一次读取会因为没有数据而阻塞
			// 后续从 client 发送来的数据会继续处理，不会引发 io.EOF
			if err != io.EOF {
				log.Fatalln(err)
			}

			// client 执行 conn.Close() 正常关闭连接时，此处的err为EOF
			log.Println(err)
			log.Printf("%s disconnect .", c.RemoteAddr())
			break
		}
		fmt.Println(buf[:n])
		all = append(all, buf[:n]...)
	}
	fmt.Println(all)
}
