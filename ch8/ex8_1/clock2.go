package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

var zones = []string{"Europe/Berlin", "Asia/Tokyo", "America/New_York", "Europe/London"}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: clock 8080 8081 ...")
	}

	for i, port := range os.Args[1:] {
		loc, _ := time.LoadLocation(zones[i%4])
		go serveTime(port, loc)
	}

	select {}
}

func serveTime(port string, loc *time.Location) {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, loc)
	}
}

var mutex = sync.Mutex{}

func handleConn(c net.Conn, loc *time.Location) {
	defer c.Close()
	for {
		mutex.Lock()
		_, err := io.WriteString(c, time.Now().In(loc).Format(fmt.Sprintf("15:04:05")))
		mutex.Unlock()
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
