// calculator/calculator.go
package calculator

import (
	"calculator_golang/mathops"
	_ "errors"
)

func Add(a, b int) int {
	return mathops.Add(a, b)
}

func Subtract(a, b int) int {
	return mathops.Subtract(a, b)
}

func Multiply(a, b int) int {
	return mathops.Multiply(a, b)
}

func Divide(a, b int) (int, error) {
	return mathops.Divide(a, b)
}

func Power(a, b int) int {
	return int(power(float64(a), float64(b)))
}

func power(a, b float64) float64 {
	result := 1.0
	for i := 0; i < int(b); i++ {
		result *= a
	}
	return result
}
