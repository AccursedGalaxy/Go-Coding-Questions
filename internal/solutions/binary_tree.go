package solutions

/*
Binary Tree Implementation

Key Concepts:
- Binary Search Tree (BST) property: Left nodes are smaller, right nodes are larger
- Recursive operations for tree traversal and manipulation
- Memory efficiency using pointer-based structure
- Time complexity: O(log n) average case for balanced trees, O(n) worst case
- Space complexity: O(h) for recursive calls, where h is tree height

This implementation provides:
1. Basic BST operations (Insert, Find)
2. Tree traversal (In-order)
3. Tree analysis (Height calculation)
4. Proper nil handling for empty trees/nodes
*/

type Node struct {
    Value int    // The value stored in this node
    Left  *Node  // Pointer to left child (smaller values)
    Right *Node  // Pointer to right child (larger values)
}

type BinaryTree struct {
    Root *Node  // Root node of the tree
}

// Insert adds a new value to the binary tree
// Time Complexity: O(log n) average case, O(n) worst case
// Space Complexity: O(h) for recursive stack
func (t *BinaryTree) Insert(value int) {
    // Special case: empty tree
    if t.Root == nil {
        t.Root = &Node{Value: value}
        return
    }
    t.insertRecursive(t.Root, value)
}

// insertRecursive is a helper function that recursively finds the correct position
// and inserts the new value while maintaining BST properties
func (t *BinaryTree) insertRecursive(node *Node, value int) {
    // Insert to left subtree if value is less than or equal
    if value <= node.Value {
        if node.Left == nil {
            // Found insertion point
            node.Left = &Node{Value: value}
        } else {
            // Continue searching in left subtree
            t.insertRecursive(node.Left, value)
        }
    } else {
        // Insert to right subtree if value is greater
        if node.Right == nil {
            // Found insertion point
            node.Right = &Node{Value: value}
        } else {
            // Continue searching in right subtree
            t.insertRecursive(node.Right, value)
        }
    }
}

// Find searches for a value in the binary tree
// Time Complexity: O(log n) average case, O(n) worst case
// Space Complexity: O(h) for recursive stack
// Returns: true if found, false otherwise
func (t *BinaryTree) Find(value int) bool {
    return t.findRecursive(t.Root, value)
}

// findRecursive is a helper function that recursively searches for the value
// using BST properties to optimize the search path
func (t *BinaryTree) findRecursive(node *Node, value int) bool {
    // Base case: reached nil node (value not found)
    if node == nil {
        return false
    }
    // Base case: found the value
    if node.Value == value {
        return true
    }
    // Recursive case: search left or right subtree based on value
    if value < node.Value {
        return t.findRecursive(node.Left, value)
    }
    return t.findRecursive(node.Right, value)
}

// InOrderTraversal returns the elements of the tree in sorted order
// Time Complexity: O(n) where n is number of nodes
// Space Complexity: O(n) for result slice + O(h) for recursive stack
func (t *BinaryTree) InOrderTraversal() []int {
    result := make([]int, 0) // Create slice with zero initial capacity
    t.inOrderRecursive(t.Root, &result)
    return result
}

// inOrderRecursive performs in-order traversal (Left-Root-Right)
// Uses pointer to slice to avoid copying and improve efficiency
func (t *BinaryTree) inOrderRecursive(node *Node, result *[]int) {
    if node != nil {
        // Process left subtree
        t.inOrderRecursive(node.Left, result)
        // Process current node
        *result = append(*result, node.Value)
        // Process right subtree
        t.inOrderRecursive(node.Right, result)
    }
}

// Height calculates the maximum depth of the tree
// Time Complexity: O(n) where n is number of nodes
// Space Complexity: O(h) for recursive stack
func (t *BinaryTree) Height() int {
    return t.heightRecursive(t.Root)
}

// heightRecursive calculates height using recursive depth-first search
// Height is defined as the number of edges in longest path from root to leaf
func (t *BinaryTree) heightRecursive(node *Node) int {
    // Base case: empty subtree has height -1
    // This ensures leaf nodes have height 0
    if node == nil {
        return -1
    }

    // Calculate height of left and right subtrees
    leftHeight := t.heightRecursive(node.Left)
    rightHeight := t.heightRecursive(node.Right)

    // Return maximum height of subtrees plus 1 for current level
    return max(leftHeight, rightHeight) + 1
}

// max is a helper function to find the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
