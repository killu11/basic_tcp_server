package main

import (
	"log"
	"net"
)

func main() {
	Start()
}

func Start() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Server already running")
	}
	defer ln.Close()
	log.Println("Server started on port `:8080`")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Bad connection: %v", err.Error())
			continue
		}
		go connHandler(conn)
	}
}

func connHandler(conn net.Conn) {
	_, err := conn.Write([]byte("OK\n"))

	if err != nil {
		log.Println(err.Error())
	}
}
