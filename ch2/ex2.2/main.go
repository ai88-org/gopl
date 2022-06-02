package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		input := bufio.NewScanner(os.Stdin)

		for input.Scan() {
			if input.Text() == "." {
				break
			}
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex2.2: %v\n", err)
				os.Exit(1)
			}
			d := Dollar(t)
			r := Renminbi(t)
			fmt.Printf("%s = %s，%s = %s\n", d, DToR(d), r, RToD(r))
		}
	} else {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			d := Dollar(t)
			r := Renminbi(t)
			fmt.Printf("%s = %s，%s = %s\n", d, DToR(d), r, RToD(r))
		}
	}
}
