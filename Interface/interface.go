// ui/interface.go
package ui

import (
	"calculator_golang/calculator"
	"errors"
	"fmt"
	_ "strconv"
)

func StartCalculatorUI() {
	var num1, num2 float64
	var operation string

	// Get user input
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Print("Choose operation (+, -, *, /, ^): ")
	fmt.Scanln(&operation)

	// Perform the calculation
	result, err := calculate(int(num1), int(num2), operation)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func calculate(a, b int, operation string) (int, error) {
	switch operation {
	case "+":
		return calculator.Add(a, b), nil
	case "-":
		return calculator.Subtract(a, b), nil
	case "*":
		return calculator.Multiply(a, b), nil
	case "/":
		return calculator.Divide(a, b)
	case "^":
		return calculator.Power(a, b), nil
	default:
		return 0, errors.New("invalid operation")
	}
}
