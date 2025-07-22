package main

import (
	"log"
	"net"
)

func main() {
	Ping()
}

func Ping() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatalln("Bad connection to `localhost:8080`")
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("Connection close!")
		}
	}(conn)

	buff := make([]byte, 1024)
	n, err := conn.Read(buff)

	if err != nil {
		log.Fatalln(err.Error())
	}

	msg := string(buff[:n])

	if msg == "OK\n" {
		log.Println("Response `OK` successfully got!\nClosing connection...")
		conn.Close()
	}
}
