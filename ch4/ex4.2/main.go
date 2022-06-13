package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var n384 = flag.Bool("n384", false, "output sha384")
var n512 = flag.Bool("n512", false, "output sha512")

func main() {
	flag.Parse()

	if *n384 {
		fmt.Println("here384")
		c1 := sha512.Sum384([]byte(os.Args[1]))
		c2 := sha512.Sum384([]byte(os.Args[2]))
		fmt.Printf("%x\n%x", c1, c2)
		return
	}

	if *n512 {
		fmt.Println("here512")
		c1 := sha512.Sum512([]byte(os.Args[1]))
		c2 := sha512.Sum512([]byte(os.Args[2]))
		fmt.Printf("%x\n%x", c1, c2)
		return
	}

	fmt.Println("here256")
	c1 := sha256.Sum256([]byte(os.Args[1]))
	c2 := sha256.Sum256([]byte(os.Args[2]))
	fmt.Printf("%x\n%x", c1, c2)

}
