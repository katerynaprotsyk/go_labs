package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{5, 27, 34, 88, 2, 56, 1, 67, 90, 12}
	result := make([]int, 10)
	for i := 0; i < 10; i++ {
		result[i] = a[i] + b[i]
	}
	fmt.Println("Результат:", result)
}
