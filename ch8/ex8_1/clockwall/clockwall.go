package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: clockwall NewYork=localhost:8010 Tokyo=localhost:8020 ...")
	}

	smap = make(map[int]string, len(os.Args)-1)

	for i, arg := range os.Args[1:] {
		pair := strings.Split(arg, "=")
		go fetchTime(pair[1], i)
	}

	select {}
}

func fetchTime(address string, index int) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	if _, err := io.Copy(&MyWriter{Index: index}, conn); err != nil {
		log.Fatal(err)
	}
}

// MyWriter ...
type MyWriter struct {
	Index int
}

var mutex = sync.RWMutex{}

var smap map[int]string

// Write ...
func (w *MyWriter) Write(b []byte) (int, error) {
	mutex.Lock()
	smap[w.Index] = string(b)
	s := fmt.Sprintf("\r%10s\t%10s\t%10s", smap[0], smap[1], smap[2])
	mutex.Unlock()
	if _, err := io.WriteString(os.Stdout, s); err != nil {
		return len(b), err
	}
	return len(b), nil
}
