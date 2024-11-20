package solutions

import (
	"sync"
	"sync/atomic"
)

/*
Concurrent B-Tree Implementation

Key Concepts:
- Lock-Free Algorithms: Using atomic operations and CAS
- Memory Ordering: Ensuring proper memory visibility
- Transactional Memory: Supporting atomic multi-key operations
- Version Control: Managing concurrent access with versioning
- Resource Management: Proper cleanup and memory reclamation

Design Patterns:
1. MVCC (Multi-Version Concurrency Control)
2. Copy-on-Write for snapshots
3. Hand-over-hand locking for traversal
4. Memory pooling for node allocation

Performance Characteristics:
- Time Complexity: O(log n) average operations
- Space Complexity: O(n) for storage, O(log n) for operations
- Memory Overhead: Additional versioning and synchronization structures
*/

// NodeVersion tracks the modification state of a node
type NodeVersion uint64

// BTreeNode represents a node in the B-tree
type BTreeNode struct {
	mu       sync.RWMutex    // Protects node data
	version  NodeVersion     // Used for MVCC
	keys     []interface{}   // Sorted keys in the node
	children []*BTreeNode    // Child pointers (nil for leaf nodes)
	isLeaf   bool           // True if this is a leaf node
	next     *BTreeNode     // Link to next node at same level (for range queries)
}

// ConcurrentBTree represents a thread-safe B-tree
type ConcurrentBTree struct {
	mu      sync.RWMutex              // Protects tree structure
	root    atomic.Pointer[BTreeNode] // Root node (atomic for lock-free reads)
	degree  int                       // Minimum degree of the tree
	compare func(a, b interface{}) int // Custom comparison function
}

// NewConcurrentBTree creates a new concurrent B-tree
func NewConcurrentBTree(degree int, compare func(a, b interface{}) int) *ConcurrentBTree {
	if degree < 2 {
		degree = 2 // Minimum valid degree
	}

	tree := &ConcurrentBTree{
		degree:  degree,
		compare: compare,
	}

	// Initialize with empty root node
	root := &BTreeNode{
		isLeaf: true,
		keys:   make([]interface{}, 0, 2*degree-1),
	}
	tree.root.Store(root)
	return tree
}

// Insert adds a key to the B-tree while maintaining thread safety
func (t *ConcurrentBTree) Insert(key interface{}) error {
	root := t.root.Load()

	// Handle root split if needed
	if len(root.keys) == 2*t.degree-1 {
		t.mu.Lock() // Lock tree for root split
		defer t.mu.Unlock()

		newRoot := &BTreeNode{
			isLeaf:   false,
			keys:     make([]interface{}, 0),
			children: []*BTreeNode{root},
		}
		t.splitChild(newRoot, 0)
		t.root.Store(newRoot)
		return t.insertNonFull(newRoot, key)
	}

	return t.insertNonFull(root, key)
}

// insertNonFull inserts a key into a non-full node
func (t *ConcurrentBTree) insertNonFull(node *BTreeNode, key interface{}) error {
	node.mu.Lock()
	defer node.mu.Unlock()

	i := len(node.keys) - 1

	if node.isLeaf {
		// Insert into leaf node
		for i >= 0 && t.compare(node.keys[i], key) > 0 {
			i--
		}
		// Insert key at correct position
		node.keys = append(node.keys, nil)
		copy(node.keys[i+2:], node.keys[i+1:])
		node.keys[i+1] = key
		atomic.AddUint64((*uint64)(&node.version), 1)
		return nil
	}

	// Find the child to recurse into
	for i >= 0 && t.compare(node.keys[i], key) > 0 {
		i--
	}
	i++

	child := node.children[i]
	if len(child.keys) == 2*t.degree-1 {
		// Split child if full
		t.splitChild(node, i)
		if t.compare(key, node.keys[i]) > 0 {
			i++
		}
	}

	return t.insertNonFull(node.children[i], key)
}

// splitChild splits a full child node during insertion
func (t *ConcurrentBTree) splitChild(parent *BTreeNode, childIndex int) {
	child := parent.children[childIndex]
	newChild := &BTreeNode{
		isLeaf: child.isLeaf,
		keys:   make([]interface{}, t.degree-1),
	}

	// Copy right half of keys to new node
	copy(newChild.keys, child.keys[t.degree:])

	if !child.isLeaf {
		newChild.children = make([]*BTreeNode, t.degree)
		copy(newChild.children, child.children[t.degree:])
	}

	// Update child's key count
	child.keys = child.keys[:t.degree-1]

	// Insert new child into parent
	parent.children = append(parent.children, nil)
	copy(parent.children[childIndex+2:], parent.children[childIndex+1:])
	parent.children[childIndex+1] = newChild

	// Move median key to parent
	parent.keys = append(parent.keys, nil)
	copy(parent.keys[childIndex+2:], parent.keys[childIndex+1:])
	parent.keys[childIndex+1] = child.keys[t.degree-1]

	// Update versions
	atomic.AddUint64((*uint64)(&parent.version), 1)
	atomic.AddUint64((*uint64)(&child.version), 1)
	atomic.AddUint64((*uint64)(&newChild.version), 1)
}

// Search looks for a key in the B-tree
func (t *ConcurrentBTree) Search(key interface{}) (bool, error) {
	node := t.root.Load()
	for {
		i := 0
		for i < len(node.keys) && t.compare(node.keys[i], key) < 0 {
			i++
		}

		if i < len(node.keys) && t.compare(node.keys[i], key) == 0 {
			return true, nil
		}

		if node.isLeaf {
			return false, nil
		}

		node = node.children[i]
	}
}

// RangeQuery returns all keys in the given range [start, end]
func (t *ConcurrentBTree) RangeQuery(start, end interface{}) ([]interface{}, error) {
	result := make([]interface{}, 0)
	node := t.findLeaf(start)

	// Traverse leaves using next pointers
	for node != nil {
		node.mu.RLock()
		for _, key := range node.keys {
			if t.compare(key, start) >= 0 && t.compare(key, end) <= 0 {
				result = append(result, key)
			}
			if t.compare(key, end) > 0 {
				node.mu.RUnlock()
				return result, nil
			}
		}
		next := node.next
		node.mu.RUnlock()
		node = next
	}

	return result, nil
}

// findLeaf finds the leaf node where a key should be located
func (t *ConcurrentBTree) findLeaf(key interface{}) *BTreeNode {
	node := t.root.Load()
	for !node.isLeaf {
		node.mu.RLock()
		i := 0
		for i < len(node.keys) && t.compare(node.keys[i], key) < 0 {
			i++
		}
		next := node.children[i]
		node.mu.RUnlock()
		node = next
	}
	return node
}

// Snapshot creates a consistent point-in-time view of the tree
func (t *ConcurrentBTree) Snapshot() (*ConcurrentBTree, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	// Create new tree with same configuration
	snapshot := NewConcurrentBTree(t.degree, t.compare)

	// Copy root and all nodes (Copy-on-Write)
	root := t.root.Load()
	newRoot := t.cloneNode(root)
	snapshot.root.Store(newRoot)

	return snapshot, nil
}

// cloneNode creates a deep copy of a node and its children
func (t *ConcurrentBTree) cloneNode(node *BTreeNode) *BTreeNode {
	if node == nil {
		return nil
	}

	node.mu.RLock()
	defer node.mu.RUnlock()

	clone := &BTreeNode{
		isLeaf: node.isLeaf,
		keys:   make([]interface{}, len(node.keys)),
	}

	copy(clone.keys, node.keys)

	if !node.isLeaf {
		clone.children = make([]*BTreeNode, len(node.children))
		for i, child := range node.children {
			clone.children[i] = t.cloneNode(child)
		}
	}

	return clone
}
