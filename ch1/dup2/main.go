package main

import (
	"bufio"
	"fmt"
	"os"
)

// 切片传入的时候有点意思
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
		}
	}

}

func countLines(file *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
}
