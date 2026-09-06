package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// Strings in Go are immutable byte sequences.
	myString := "resume"
	message := "Go is awesome"
	unicodeWord := "Golang-你好"

	fmt.Println("String:", myString)
	fmt.Println("Upper:", strings.ToUpper(myString))
	fmt.Println("Contains 'awesome':", strings.Contains(message, "awesome"))
	fmt.Println("Replace:", strings.ReplaceAll(message, "awesome", "powerful"))

	// len returns bytes, not Unicode characters.
	fmt.Println("Byte length:", len(unicodeWord))
	// RuneCountInString counts Unicode code points.
	fmt.Println("Rune count:", utf8.RuneCountInString(unicodeWord))

	// range over a string decodes runes.
	for index, char := range unicodeWord {
		fmt.Printf("index=%d char=%c\n", index, char)
	}
}
