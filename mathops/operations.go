package mathops

import "errors"

// Add performs addition of two numbers.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract performs subtraction of two numbers.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply performs multiplication of two numbers.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide performs division of two numbers.
// It returns both the quotient and remainder.
// It returns an error if the divisor is zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
