package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	url := os.Args[1]
	fileName := os.Args[2]
	go fecth(url, fileName, ch)
	// 这里为啥要遍历的从管道取数呢
	fmt.Println(<-ch)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fecth(url string, fileName string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintln("这里就报错了", err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	// 这里忽视一下错误的处理 ignore write error
	ioutil.WriteFile(fileName, b, 0777)
	if err != nil {
		ch <- fmt.Sprintln(err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%s", secs, url)
}
