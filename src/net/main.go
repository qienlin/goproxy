package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("Error listening on port 8000: ", err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close()
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			break
		}
		fmt.Println(string(buf[0:n]))
	}
}
