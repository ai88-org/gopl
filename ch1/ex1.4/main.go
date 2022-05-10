package main

import (
	"bufio"
	"fmt"
	"os"
)

// 切片传入的时候有点意思
func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "os.Stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}

	for fileName, map1 := range counts {
		for line, n := range map1 {
			if n > 1 {
				fmt.Printf("%s\t%s\t%d\n", fileName, line, n)
			}
		}

	}
}

func countLines(file *os.File, counts map[string]map[string]int, fileName string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counts[fileName][scanner.Text()]++
	}
}
