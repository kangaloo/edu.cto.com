package main

import (
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

	input := ""

	for {
		_, err = fmt.Scanf("%v", &input)
		if err != nil {
			conn.Close()
			log.Fatalln(err)
		}

		if input == "exit" {
			conn.Close()
			os.Exit(0)
		}

		conn.Write([]byte(input))
	}

}
