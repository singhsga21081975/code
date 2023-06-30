package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a parent context
	parentCtx := context.Background()

	// Create a derived context with cancel function
	ctx, cancel := context.WithCancel(parentCtx)

	// Start a goroutine that performs some task
	go task(ctx)

	// Wait for a few seconds
	time.Sleep(5 * time.Second)

	// Cancel the context to stop the task
	cancel()

	// Wait for a moment to allow task to stop gracefully
	time.Sleep(1 * time.Second)

	fmt.Println("Program finished")
}

func task(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// Context has been cancelled, stop the task
			fmt.Println("Task stopped")
			return
		default:
			// Task is still running
			fmt.Println("Task running")
			time.Sleep(1 * time.Second)
		}
	}
}
