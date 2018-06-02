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
			log.Fatalln(err)
		}

		if input == "exit" {
			os.Exit(0)
		}

		if input == "" {
			continue
		}

		conn.Write([]byte(input))
	}
}
