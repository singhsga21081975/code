package main

import (
	"context"
	"fmt"
	"time"
)

type key string

func main() {
	// Create a parent context with cancellation capability
	parentCtx, cancelParent := context.WithCancel(context.Background())

	// Create a child context for the first goroutine with a value
	childCtx1 := context.WithValue(parentCtx, key("name"), "Alice")

	// Start the first goroutine with the child context
	go doWork(childCtx1)

	// Create a child context for the second goroutine with a value
	childCtx2 := context.WithValue(parentCtx, key("name"), "Bob")

	// Start the second goroutine with the child context
	go doWork(childCtx2)

	// Simulate a delay before canceling the parent context
	time.Sleep(3 * time.Second)

	// Cancel the parent context to signal termination to all child goroutines
	cancelParent()

	// Wait for the child goroutines to finish
	time.Sleep(1 * time.Second)

	fmt.Println("Program terminated")
}

func doWork(ctx context.Context) {
	// Get the value from the context
	name, ok := ctx.Value(key("name")).(string)
	if !ok {
		name = "Unknown"
	}

	for {
		select {
		case <-ctx.Done():
			// Context canceled, terminate the goroutine
			fmt.Printf("%s: Work canceled\n", name)
			return
		default:
			// Perform some work
			fmt.Printf("%s: Doing work...\n", name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
