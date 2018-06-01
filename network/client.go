package main

import (
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")

	if err != nil {
		log.Fatalln(err)
	}

	conn.Write([]byte("abc"))

}
