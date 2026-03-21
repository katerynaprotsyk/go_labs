package calc

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("Пакет calc ініціалізовано")
}

func Sum(nums ...float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total
}

func Max(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for _, n := range nums {
		if n > res {
			res = n
		}
	}
	return res
}

func Min(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for _, n := range nums {
		if n < res {
			res = n
		}
	}
	return res
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("ділення на нуль неможливе")
	}
	return a / b, nil
}
