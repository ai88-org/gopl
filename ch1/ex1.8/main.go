package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// 增加对前缀的判断
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)

		if err != nil {
			os.Exit(1)
		}
		defer resp.Body.Close()
		b, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			os.Exit(2)
		}
		fmt.Printf("%s\n", b)
	}

}
