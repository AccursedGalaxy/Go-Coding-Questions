package solutions

import "fmt"

/*
Division with Error Handling

Key Concepts:
- Error handling in Go
- Custom error types
- Multiple return values
- Mathematical edge cases

This implementation demonstrates proper error handling in Go by:
1. Using a custom error type for specific error cases
2. Implementing the error interface
3. Providing detailed error messages
4. Handling edge cases (division by zero)
*/

// DivisionError is a custom error type that provides context about division errors
// It implements the error interface through its Error() method
type DivisionError struct {
    dividend float64 // The number being divided (numerator)
    divisor  float64 // The number dividing by (denominator)
    message  string  // A human-readable error message
}

// Error implements the error interface for DivisionError
// This method formats the error message with specific details about the failed operation
func (e *DivisionError) Error() string {
    // fmt.Sprintf creates a formatted string containing:
    // - The dividend (the number being divided)
    // - The divisor (the number dividing by)
    // - The specific error message
    // %v is used for default formatting of the float64 values
    return fmt.Sprintf("cannot divide %v by %v: %s", e.dividend, e.divisor, e.message)
}

// Divide performs division of two floating-point numbers
// Parameters:
//   - a: dividend (numerator)
//   - b: divisor (denominator)
// Returns:
//   - float64: the result of a/b
//   - error: nil if successful, DivisionError if division by zero
func Divide(a, b float64) (float64, error) {
    // Check for division by zero
    // This is a critical mathematical error case that must be handled
    if b == 0 {
        // Return 0 and a new DivisionError with context about the failed operation
        return 0, &DivisionError{
            dividend: a,
            divisor:  b,
            message:  "division by zero",
        }
    }
    
    // If divisor is not zero, perform the division and return the result
    // No error is returned (nil) since the operation was successful
    return a / b, nil
} 