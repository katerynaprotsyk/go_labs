package main

import "fmt"

func generate() <-chan int {
	c := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func filterEven(in <-chan int) <-chan int {
	c := make(chan int, 10)
	go func() {
		for n := range in {
			if n%2 == 0 {
				c <- n
			}
		}
		close(c)
	}()
	return c
}

func square(in <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for n := range in {
			c <- n * n
		}
		close(c)
	}()
	return c
}

func sum(in <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		s := 0
		for n := range in {
			s += n
		}
		c <- s
		close(c)
	}()
	return c
}

func main() {
	g := generate()
	f := filterEven(g)
	q := square(f)
	s := sum(q)

	res := <-s
	fmt.Println("Результат:", res)
}
