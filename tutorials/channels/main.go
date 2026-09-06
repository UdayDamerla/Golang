package main

import (
	"fmt"
	"time"
)

// Channels in Go are used to send and receive values between goroutines.
// This tutorial shows common channel patterns used in real code.

func unbufferedChannelDemo() {
	fmt.Println("=== 1) Unbuffered channel ===")
	myChannel := make(chan string)

	go func() {
		myChannel <- "Hello, World!"
	}()

	message := <-myChannel
	fmt.Println("Received:", message)
}

func bufferedChannelDemo() {
	fmt.Println("\n=== 2) Buffered channel ===")
	buffered := make(chan int, 3)

	buffered <- 10
	buffered <- 20
	buffered <- 30

	fmt.Println("Buffered values:", <-buffered, <-buffered, <-buffered)
}

func closeChannelAndRangeDemo() {
	fmt.Println("\n=== 3) Close channel + range ===")
	numbers := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	for n := range numbers {
		fmt.Println("Range received:", n)
	}
}

func receiveWithOkDemo() {
	fmt.Println("\n=== 4) Receive with ok check ===")
	ch := make(chan string, 1)
	ch <- "Go Channels"
	close(ch)

	v1, ok1 := <-ch
	fmt.Println("First receive:", v1, "ok:", ok1)

	v2, ok2 := <-ch
	fmt.Println("Second receive after close:", v2, "ok:", ok2)
}

func selectDemo() {
	fmt.Println("\n=== 5) select with timeout ===")
	dataCh := make(chan string)

	go func() {
		time.Sleep(150 * time.Millisecond)
		dataCh <- "Data loaded"
	}()

	select {
	case msg := <-dataCh:
		fmt.Println("select received:", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timeout while waiting for data")
	}
}

func processData(data string) {
	fmt.Println("Processing data:", data)
}

func worker(input <-chan string, done chan<- bool) {
	for item := range input {
		processData(item)
	}
	done <- true
}

func directionChannelDemo() {
	fmt.Println("\n=== 6) Directional channels (send-only / receive-only) ===")
	input := make(chan string)
	done := make(chan bool)

	go worker(input, done)

	input <- "task-1"
	input <- "task-2"
	close(input)

	<-done
	fmt.Println("Worker finished")
}

func main() {
	unbufferedChannelDemo()
	bufferedChannelDemo()
	closeChannelAndRangeDemo()
	receiveWithOkDemo()
	selectDemo()
	directionChannelDemo()
}
