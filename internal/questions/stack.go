package questions

/*
Challenge: Data Structures

Problem: Implement a stack data structure with the following methods:
1. Push: Add an element to the top of the stack.
2. Pop: Remove the top element from the stack.
3. Peek: Return the top element without removing it.
4. IsEmpty: Check if the stack is empty.

Requirement: Use a slice for the underlying data structure.
*/

type Stack struct {
    elements []int
}

func (s *Stack) Push(value int)     {}
func (s *Stack) Pop() (int, error)  { return 0, nil }
func (s *Stack) Peek() (int, error) { return 0, nil }
func (s *Stack) IsEmpty() bool      { return true }
