package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
)

var mu sync.Mutex
var count int

// 这个例子使用浏览器去调试会出现其他结果哦
// 使用fetch下面的代码调试
func main() {
	runtime.GOMAXPROCS(1)
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count = count + 1
	mu.Unlock()
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
