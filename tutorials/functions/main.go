package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Functions in Go")

	// Go functions can return multiple values.
	quotient, remainder := divide(10, 3)
	fmt.Printf("10 / 3 => quotient=%d remainder=%d\n", quotient, remainder)

	// Errors are returned as regular values in Go.
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("safeDivide error:", err)
	} else {
		fmt.Println("safeDivide result:", result)
	}

	fmt.Println("sum(1,2,3,4,5) =", sum(1, 2, 3, 4, 5))
}

// Functions use typed parameters and explicit return types.
func printMe(printValue string) {
	fmt.Println(printValue)
}

// This returns quotient and remainder.
func divide(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}

// Returning error is the standard Go way to report failure.
func safeDivide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return num1 / num2, nil
}

// Variadic parameters accept zero or more arguments.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
