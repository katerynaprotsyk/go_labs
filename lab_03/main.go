package main

import (
	"awesomeProject/lab_03/calc"
	"fmt"
)

type Calculator interface {
	Sum(nums ...float64) float64
	Max(nums ...float64) float64
	Min(nums ...float64) float64
	Divide(a, b float64) (float64, error)
}

type Calc struct{}

func (c Calc) Sum(nums ...float64) float64          { return calc.Sum(nums...) }
func (c Calc) Max(nums ...float64) float64          { return calc.Max(nums...) }
func (c Calc) Min(nums ...float64) float64          { return calc.Min(nums...) }
func (c Calc) Divide(a, b float64) (float64, error) { return calc.Divide(a, b) }

func main() {
	sumRes := calc.Sum(10, 20, 30.5)
	maxRes := calc.Max(1.2, 5.9, 2.3)
	minRes := calc.Min(1.2, 5.9, 2.3)

	fmt.Printf("Сума: %v, Макс: %v, Мін: %v\n", sumRes, maxRes, minRes)

	res, err := calc.Divide(10, 15)
	if err != nil {
		fmt.Println("Помилка:", err)
		return
	}
	fmt.Println("Ділення:", res)

	fmt.Println("Завдання 2")

	var calculator Calculator = Calc{}

	fmt.Println("Sum:", calculator.Sum(1, 2, 3, 4))
	fmt.Println("Max:", calculator.Max(12, 458, 2, 832))
	fmt.Println("Min:", calculator.Min(10, 582, 310, 738))

	result, err := calculator.Divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Divide:", result)
	}

}
