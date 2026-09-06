package main

import (
	"fmt"
)

func main() {
	// Basic output
	fmt.Println("Hello, Go!")

	// Variable declarations
	var age int = 30
	var city = "Bengaluru"
	// := lets Go infer the type inside functions.
	country := "India"

	// Constants
	const appName string = "Go Basics"

	// Type conversion example
	var score int = 95
	percentage := float64(score) / 100.0

	// Multiple variable declaration
	var a, b int = 10, 20

	// fmt.Println adds spaces between arguments automatically.
	fmt.Println("App:", appName)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Country:", country)
	fmt.Println("Percentage:", percentage)
	fmt.Println("a + b =", a+b)
}
