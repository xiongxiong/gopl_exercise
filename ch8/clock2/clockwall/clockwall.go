package main

import (
	"log"
	"net"
)

func main() {

}

func fetchTime(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	if _, err :
}
