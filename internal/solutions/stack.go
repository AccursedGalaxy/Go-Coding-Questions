package solutions

import "errors"

/*
Stack Implementation

Key Concepts:
- LIFO (Last In, First Out) data structure
- Slice-based implementation for dynamic sizing
- Error handling for edge cases
- Method receiver functions
- Zero-value initialization

This implementation provides:
1. Standard stack operations (Push, Pop, Peek)
2. Error handling for empty stack operations
3. Efficient memory usage with slices
4. Thread-unsafe operations (for simplicity)
*/

// Stack represents a LIFO (Last In, First Out) data structure
// The zero value for Stack is a valid empty stack
type Stack struct {
    elements []int // Underlying slice to store stack elements
}

// Push adds a new element to the top of the stack
// Parameters:
//   - value: the integer to add to the stack
// Time complexity: O(1) amortized
func (s *Stack) Push(value int) {
    // append automatically handles growing the underlying array when needed
    // The new element is added to the end of the slice, which represents the top of the stack
    s.elements = append(s.elements, value)
}

// Pop removes and returns the top element from the stack
// Returns:
//   - int: the top element
//   - error: nil if successful, error if stack is empty
// Time complexity: O(1)
func (s *Stack) Pop() (int, error) {
    // Check if the stack is empty before attempting to pop
    if s.IsEmpty() {
        return 0, errors.New("stack is empty")
    }

    // Get the last element (top of stack)
    value := s.elements[len(s.elements)-1]

    // Remove the last element by slicing
    // This maintains the underlying array but reduces the length
    s.elements = s.elements[:len(s.elements)-1]

    return value, nil
}

// Peek returns the top element without removing it
// Returns:
//   - int: the top element
//   - error: nil if successful, error if stack is empty
// Time complexity: O(1)
func (s *Stack) Peek() (int, error) {
    // Check if the stack is empty before attempting to peek
    if s.IsEmpty() {
        return 0, errors.New("stack is empty")
    }

    // Return the last element without modifying the stack
    return s.elements[len(s.elements)-1], nil
}

// IsEmpty checks if the stack is empty
// Returns:
//   - bool: true if the stack is empty, false otherwise
// Time complexity: O(1)
func (s *Stack) IsEmpty() bool {
    return len(s.elements) == 0
}
