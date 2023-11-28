package calculator

import (
	"calculator_golang/mathops"
	_ "errors"
)

// Add performs addition of two numbers using the mathops package.
func Add(a, b float64) float64 {
	return mathops.Add(a, b)
}

// Subtract performs subtraction of two numbers using the mathops package.
func Subtract(a, b float64) float64 {
	return mathops.Subtract(a, b)
}

// Multiply performs multiplication of two numbers using the mathops package.
func Multiply(a, b float64) float64 {
	return mathops.Multiply(a, b)
}

// Divide performs division of two numbers using the mathops package.
// It returns both the quotient and an error if the divisor is zero.
func Divide(a, b float64) (float64, error) {
	return mathops.Divide(a, b)
}

// Power calculates the result of raising a to the power of b.
// It uses the power function, and handles negative exponents appropriately.
func Power(a, b float64) float64 {
	if b >= 0 {
		return power(a, b)
	} else {
		// Handle negative exponent
		return 1.0 / power(a, -b)
	}
}

// power calculates the result of raising a to the power of b.
// It uses a simple iterative approach by multiplying a by itself b times.
// This function assumes b is a non-negative integer.
func power(a, b float64) float64 {
	result := 1.0

	// Iterate b times and multiply 'result' by 'a' in each iteration
	for i := 0; i < int(b); i++ {
		result *= a
	}
	return result
}
