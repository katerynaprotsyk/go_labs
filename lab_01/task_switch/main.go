package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var y int32
	x := rand.Int31n(500)

	switch {
	case x > 10 && x < 100:
		y = 100 - (x + x*x)
	default:
		y = x - 15*(x*x)
	}

	fmt.Printf("x: %d\n", x)
	fmt.Printf("Result y: %d\n", y)
}
