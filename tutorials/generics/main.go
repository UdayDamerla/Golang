package main

import (
	"fmt"
)

// Generics in Go let us write reusable, type-safe code.
// This file demonstrates common generic patterns used in real projects.

// Number is a type constraint.
// Only these numeric types can be used where Number is required.
type Number interface {
	int | int64 | float64
}

// PrintValue accepts any type.
func PrintValue[T any](value T) {
	fmt.Println("Value:", value)
}

// Sum adds all values in a slice of numeric types.
func Sum[T Number](values []T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}

// Reverse returns a reversed copy of a slice.
func Reverse[T any](items []T) []T {
	result := make([]T, len(items))
	for i := range items {
		result[len(items)-1-i] = items[i]
	}
	return result
}

// Contains checks whether target exists in a slice.
// comparable allows == and != operations.
func Contains[T comparable](items []T, target T) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}

// Pair is a generic struct with two potentially different types.
type Pair[A any, B any] struct {
	First  A
	Second B
}

// Stack is a generic LIFO data structure.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	lastIndex := len(s.items) - 1
	value := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return value, true
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

// MapSlice transforms each item in a slice using a generic mapper function.
func MapSlice[T any, R any](items []T, mapper func(T) R) []R {
	result := make([]R, 0, len(items))
	for _, item := range items {
		result = append(result, mapper(item))
	}
	return result
}

func main() {
	fmt.Println("=== 1) Basic generic function ===")
	PrintValue(42)
	PrintValue("Hello, Generics!")
	PrintValue(3.14)

	fmt.Println("\n=== 2) Generic sum with constraints ===")
	intValues := []int{1, 2, 3, 4, 5}
	floatValues := []float64{1.5, 2.5, 3.5}
	fmt.Println("Sum of ints:", Sum(intValues))
	fmt.Println("Sum of floats:", Sum(floatValues))

	fmt.Println("\n=== 3) Generic reverse ===")
	letters := []string{"A", "B", "C", "D"}
	fmt.Println("Original:", letters)
	fmt.Println("Reversed:", Reverse(letters))

	fmt.Println("\n=== 4) Generic contains ===")
	numbers := []int{10, 20, 30, 40}
	fmt.Println("Contains 20:", Contains(numbers, 20))
	fmt.Println("Contains 99:", Contains(numbers, 99))

	fmt.Println("\n=== 5) Generic struct ===")
	user := Pair[string, int]{First: "Alice", Second: 28}
	fmt.Println("Pair value:", user)

	fmt.Println("\n=== 6) Generic stack ===")
	var stack Stack[string]
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")
	fmt.Println("Stack size after pushes:", stack.Size())

	for {
		value, ok := stack.Pop()
		if !ok {
			break
		}
		fmt.Println("Popped:", value)
	}
	fmt.Println("Stack size after pops:", stack.Size())

	fmt.Println("\n=== 7) Generic map/transform ===")
	values := []int{1, 2, 3, 4}
	squares := MapSlice(values, func(n int) int { return n * n })
	labels := MapSlice(values, func(n int) string { return fmt.Sprintf("item-%d", n) })
	fmt.Println("Input:", values)
	fmt.Println("Squares:", squares)
	fmt.Println("Labels:", labels)
}
