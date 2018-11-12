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

var zones = []string{"CST", "MST", "PST", "AKST"}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: clock 8080 8081 ...")
	}

	for i, port := range os.Args[1:] {
		go serveTime(port, zones[i%4])
	}

	select {}
}

func serveTime(port string, zone string) {
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
		go handleConn(conn, zone)
	}
}

var mutex = sync.Mutex{}

func handleConn(c net.Conn, zone string) {
	defer c.Close()
	for {
		mutex.Lock()
		println(zone)
		println(time.Now().Format(fmt.Sprintf("15:04:05 %s %s | ", zone, os.Getenv("TZ"))))
		_, err := io.WriteString(c, time.Now().Format(fmt.Sprintf("15:04:05 %s %s | ", zone, os.Getenv("TZ"))))
		mutex.Unlock()
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
