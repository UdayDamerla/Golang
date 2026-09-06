// Pointers

package main

import (
	"fmt"
)

type Counter struct {
	Value int
}

// Passing a pointer lets the function update the original value.
func increment(n *int) {
	(*n)++
}

// Swapping works by dereferencing both pointers.
func swap(a *int, b *int) {
	*a, *b = *b, *a
}

// Struct pointers are common when mutating shared state.
func updateCounter(c *Counter, value int) {
	c.Value = value
}

func main() {
	var myInt int = 42
	var myIntPointer *int = &myInt
	// 2 roles of * syntax: pointer type declaration and dereferencing.

	fmt.Println("Value of myInt:", myInt)
	fmt.Println("Address of myInt:", &myInt)
	fmt.Println("Value of myIntPointer:", myIntPointer)
	fmt.Println("Value pointed to by myIntPointer:", *myIntPointer)

	*myIntPointer = 100
	fmt.Println("After direct pointer update:", myInt)

	increment(&myInt)
	fmt.Println("After increment function:", myInt)

	a, b := 10, 20
	swap(&a, &b)
	fmt.Println("After swap a,b:", a, b)

	counter := &Counter{Value: 1}
	updateCounter(counter, 99)
	fmt.Println("Updated struct through pointer:", counter.Value)
}
