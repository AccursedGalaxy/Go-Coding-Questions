package solutions

import (
	"context"
	"fmt"
	"time"
)

/*
Channel Communication Implementation

Key Concepts:
- Producer-Consumer Pattern: Using goroutines and channels for data flow
- Context Usage: Managing timeouts and cancellation
- Error Propagation: Proper error handling across goroutines
- Channel Directionality: Using directional channels for type safety
- Graceful Shutdown: Ensuring all resources are properly cleaned up

Architecture:
1. Generator goroutine (Producer) -> numberChan ->
2. Squarer goroutine (Processor) -> squareChan ->
3. Main goroutine (Consumer)

Error Handling:
- Timeout via context
- Channel closing signals
- Error propagation through error channel
*/

// ProcessNumbers coordinates the number processing pipeline
// Parameters:
//   - n: the upper limit of numbers to process
//   - timeout: maximum duration to wait for completion
// Returns:
//   - error: nil if successful, error if timeout or processing fails
func ProcessNumbers(n int, timeout time.Duration) error {
    // Create context with timeout for graceful shutdown
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel() // Ensure resources are cleaned up

    // Create buffered channels to prevent goroutine leaks
    numberChan := make(chan int, n)    // Numbers from generator
    squareChan := make(chan int, n)    // Squared results
    errorChan := make(chan error, 2)   // Potential errors from goroutines

    // Launch producer goroutine
    go generateNumbers(ctx, numberChan, n)

    // Launch processor goroutine
    go squareNumbers(ctx, numberChan, squareChan, errorChan)

    // Process results in main goroutine
    return processResults(ctx, squareChan, errorChan)
}

// generateNumbers produces a sequence of numbers from 1 to n
// Parameters:
//   - ctx: context for cancellation
//   - out: channel to send generated numbers
//   - n: upper limit of numbers to generate
func generateNumbers(ctx context.Context, out chan<- int, n int) {
    // Ensure channel is closed when function returns
    defer close(out)

    for i := 1; i <= n; i++ {
        select {
        case <-ctx.Done():
            // Context cancelled or timed out
            return
        case out <- i:
            // Successfully sent number
            // Continue to next number
        }
    }
}

// squareNumbers processes incoming numbers and sends their squares
// Parameters:
//   - ctx: context for cancellation
//   - in: channel receiving numbers to square
//   - out: channel to send squared results
//   - errChan: channel to send any errors
func squareNumbers(ctx context.Context, in <-chan int, out chan<- int, errChan chan<- error) {
    // Ensure channels are closed when function returns
    defer close(out)
    defer close(errChan)

    for num := range in {
        select {
        case <-ctx.Done():
            // Context cancelled or timed out
            errChan <- ctx.Err()
            return
        case out <- num * num:
            // Successfully sent squared number
            // Continue to next number
        }
    }
}

// processResults handles the results and errors from the processing pipeline
// Parameters:
//   - ctx: context for cancellation
//   - squareChan: channel receiving squared numbers
//   - errorChan: channel receiving errors
// Returns:
//   - error: nil if successful, error if processing fails
func processResults(ctx context.Context, squareChan <-chan int, errorChan <-chan error) error {
    for {
        select {
        case <-ctx.Done():
            // Operation timed out or was cancelled
            return fmt.Errorf("operation timed out: %w", ctx.Err())
        
        case err, ok := <-errorChan:
            if !ok {
                // Error channel closed, switch to nil to prevent further selects
                errorChan = nil
                continue
            }
            return fmt.Errorf("processing error: %w", err)
        
        case square, ok := <-squareChan:
            if !ok {
                // Square channel closed
                if errorChan == nil {
                    // Both channels closed, processing complete
                    return nil
                }
                // Switch to nil to prevent further selects
                squareChan = nil
                continue
            }
            fmt.Printf("Square: %d\n", square)
        }
    }
} 