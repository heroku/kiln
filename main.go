package main

import (
	"log"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	log.Println("Connection!")
	conn.Close()
}

func main() {

	ln, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err)
			continue
		}
		go handleConnection(conn)
	}

}
