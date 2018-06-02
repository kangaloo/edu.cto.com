package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)
	header := bytes.NewBuffer(make([]byte, 0))

	for {

		fmt.Printf(">> ")
		input, err := reader.ReadBytes('\n')

		// 目前认为，输入文件为 os.Stdin 所以读不到 io.EOF
		if err != nil {
			return
		}

		if len(input) == 0 {
			continue
		}

		if string(input) == "exit\n" {
			os.Exit(0)
		}

		length := int64(len(input) - 1)
		err = binary.Write(header, binary.BigEndian, length)

		if err != nil {
			log.Println(err)
			return
		}

		conn.Write(header.Bytes())
		conn.Write(input[:len(input)-1])

		_, err = conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(buf))

	}
}
