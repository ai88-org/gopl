package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	// 端口
	var port = flag.String("port", os.Args[2], "port no")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)

	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // client disconnected
		}
		if err != nil {
			return // client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
