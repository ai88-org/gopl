package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, fileName := range os.Args[1:] {
		// 把文件内容一次性的全部读入内存中，返回字节切片
		// 字节切片就是这个文件的所有内容
		contents, err := ioutil.ReadFile(fileName)
		if err != nil {
			continue
		}
		// 需要先将其转换成string类型，再进行切片处理
		for _, lines := range strings.Split(string(contents), "\n") {
			counts[lines]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
		}
	}

}
