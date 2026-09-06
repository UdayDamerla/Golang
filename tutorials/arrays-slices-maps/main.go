package main

import (
	"fmt"
)

func main() {
	// Arrays have fixed length and are part of the type.
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)
	fmt.Println("Array length:", len(arr))

	// Slices are lightweight views over arrays.
	slice := arr[1:4]
	fmt.Println("Slice from array [1:4]:", slice)
	fmt.Println("Slice len/cap:", len(slice), cap(slice))

	// append may grow the backing array when needed.
	dynamic := []int{10, 20, 30}
	fmt.Println("Dynamic slice:", dynamic, "len:", len(dynamic), "cap:", cap(dynamic))
	dynamic = append(dynamic, 40, 50)
	fmt.Println("After append:", dynamic, "len:", len(dynamic), "cap:", cap(dynamic))

	// copy duplicates elements into another slice.
	copyTarget := make([]int, len(dynamic))
	copy(copyTarget, dynamic)
	fmt.Println("Copied slice:", copyTarget)

	// Maps store key-value pairs.
	myMap := map[string]int{"adam": 1, "bob": 2, "charlie": 3}
	fmt.Println("Map:", myMap)

	// The second value reports whether the key exists.
	value, ok := myMap["adam"]
	fmt.Println("adam exists:", ok, "value:", value)

	delete(myMap, "bob")
	fmt.Println("Map after deleting bob:", myMap)

	// range iterates over map entries.
	for name, id := range myMap {
		fmt.Printf("%s -> %d\n", name, id)
	}
}
