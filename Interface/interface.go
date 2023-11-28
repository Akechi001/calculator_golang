// ui/interface.go
package ui

import (
	"calculator_golang/calculator"
	"errors"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	_ "strconv"
)

func StartWebCalculator() {
	//membuat router dengan konfigurasi default
	router := gin.Default()

	//memuat template HTML pada directory templates
	router.LoadHTMLGlob("templates/*")

	//handler untuk ke rute HTTP GET pada path "/"(root)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil) //
	})

	//menyediakan akses ke file statis
	router.Static("/static", "./static")

	// Define route for the calculator
	router.POST("/calculate", calculateHandler)

	// Start the web server
	router.Run(":8080")
}

// calculateHandler is a request handler for the "/calculate" route.
// It retrieves input values (num1, num2, operation) from the request,
// performs the calculation using the calculate function, and responds
// with the result or an error in JSON format.
func calculateHandler(c *gin.Context) {
	var num1, num2 float64
	var operation string

	// Get data from the request
	num1, _ = strconv.ParseFloat(c.PostForm("num1"), 64)
	num2, _ = strconv.ParseFloat(c.PostForm("num2"), 64)
	operation = c.PostForm("operation")

	// Perform the calculation
	result, err := calculate(float64(int(num1)), float64(int(num2)), operation)
	if err != nil {
		// Respond with an error if the calculation fails
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		// Respond with the result if the calculation is successful
		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}

// calculate performs arithmetic calculations based on the provided operation.
// It utilizes the calculator package functions for basic arithmetic operations.
// The result and any error are returned to the caller.
func calculate(a, b float64, operation string) (float64, error) {
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
		// Return an error for an invalid operation
		return 0, errors.New("invalid operation")
	}
}
