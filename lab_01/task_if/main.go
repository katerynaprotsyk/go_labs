package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var y int32
	x := rand.Int31n(500)

	if x > 10 && x < 100 {
		y = 100 - (x + x*x)
	} else {
		y = x - 15*(x*x)
	}

	fmt.Printf("x: %d\n", x)
	fmt.Printf("Result y: %d\n", y)
}
