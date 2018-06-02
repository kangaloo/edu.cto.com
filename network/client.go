package main

import (
	"bufio"
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

		conn.Write(input[:len(input)-1])
	}
}
