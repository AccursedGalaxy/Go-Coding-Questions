package questions

/*
Challenge: Binary Tree Operations

Problem: Implement a binary tree with the following operations:
1. Insert a value
2. Find a value
3. Implement in-order traversal
4. Calculate tree height

Requirements:
- Use recursive approach where appropriate
- Handle empty tree cases
- Return appropriate errors when needed
*/

type Node struct {
    Value int
    Left  *Node
    Right *Node
}

type BinaryTree struct {
    Root *Node
}

func (t *BinaryTree) Insert(value int)        {}
func (t *BinaryTree) Find(value int) bool     { return false }
func (t *BinaryTree) InOrderTraversal() []int { return nil }
func (t *BinaryTree) Height() int             { return 0 } 