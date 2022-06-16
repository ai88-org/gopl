package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	for i := 1; i <= 3; i++ {
		zoneHost := strings.Split(os.Args[i], "=")[1]
		go mustCopy(os.Stdout, zoneHost)
	}
	time.Sleep(10 * time.Hour)
}

func mustCopy(dst io.Writer, host string) {
	conn, err := net.Dial("tcp", host)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(dst, conn); err != nil {
		log.Fatal(err)
	}
}
