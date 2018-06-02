package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	defer listener.Close()

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

	// 前8个 bytes 为接收长度
	header := make([]byte, 8)
	all := make([]byte, 0, 1024)
	buf := make([]byte, 1024)
	var length int64
	var n int

	for {
		// 读完最后的数据后，会阻塞
		_, err := c.Read(header)

		if err != nil {
			if err == io.EOF {
				fmt.Println("io.EOF")
			}
			fmt.Println(err)
			log.Println("read header failed .")
			c.Close()
			return
		}

		fmt.Println("header: ", header) // test

		err = binary.Read(bytes.NewBuffer(header), binary.BigEndian, &length)

		fmt.Println("length: ", length) // test

		if err != nil {
			log.Println(err)
		}

		l := int(length)

		for {
			var err error

			fmt.Println("l: ", l)

			if l < len(buf) {
				n, err = c.Read(buf[:l])
			} else {
				n, err = c.Read(buf)
			}

			if err != nil {

				// 从socket读取完所有数据的当次，或者下一次会得到 io.EOF
				// 目前没有出现在当次，又因为当次读取完成后，下一次读取会因为没有数据而阻塞
				// 后续从 client 发送来的数据会继续处理，不会引发 io.EOF
				// todo 如何才能引发这个 io.EOF
				if err != io.EOF {
					log.Fatalln("io.EOF", err)
				}

				// client 执行 conn.Close() 正常关闭连接时，此处的err为EOF
				log.Println("EOF", err)
				log.Printf("%s disconnect .", c.RemoteAddr())
				break
			}

			fmt.Println("content: ", buf[:n])

			// todo 写文件
			all = append(all, buf[:n]...)
			l -= n

			fmt.Println("new l: ", l)

			if l <= 0 {
				c.Write([]byte("ok"))
				fmt.Println("read command finished .")
				fmt.Println()
				break
			}
		}
	}
	fmt.Println(all)
}
