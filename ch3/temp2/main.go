package main

import (
	"fmt"
	"math"
)

func main() {
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)

}
