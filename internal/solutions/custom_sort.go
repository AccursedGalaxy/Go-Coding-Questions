package solutions

import (
	"fmt"
	"sort"
	"sync"
)

/*
Custom Sort Implementation

Key Concepts:
- sort.Interface Implementation: Required methods (Len, Less, Swap)
- Thread Safety: Using sync.RWMutex for concurrent access
- Flexible Comparison: Supporting multiple fields and sort directions
- Error Handling: Validating sort criteria
- Type Safety: Strong typing for sort fields

Design Patterns:
1. Strategy Pattern: Different sorting strategies based on field
2. Mutex Pattern: Read/Write locking for thread safety
3. Builder Pattern: Flexible construction with sort criteria

Performance Characteristics:
- Time Complexity: O(n log n) from sort.Sort
- Space Complexity: O(1) additional space
- Lock Contention: Minimal due to RWMutex usage
*/

// Person represents an individual with sortable attributes
type Person struct {
	Name   string  // Person's name (string comparison)
	Age    int     // Person's age (numeric comparison)
	Height float64 // Person's height (floating-point comparison)
}

// PersonCollection represents a thread-safe collection of Person objects
// that can be sorted by different criteria
type PersonCollection struct {
	sync.RWMutex           // Embedded mutex for thread safety
	People    []Person     // Slice of Person objects to sort
	SortField string       // Field to sort by ("name", "age", "height")
	Ascending bool         // Sort direction (true = ascending, false = descending)
}

// Len implements sort.Interface
// Returns the number of elements in the collection
// Thread-safe: Uses RLock for concurrent read access
func (pc *PersonCollection) Len() int {
	pc.RLock()
	defer pc.RUnlock()
	return len(pc.People)
}

// Swap implements sort.Interface
// Swaps elements at indices i and j
// Thread-safe: Uses Lock for exclusive write access
func (pc *PersonCollection) Swap(i, j int) {
	pc.Lock()
	defer pc.Unlock()
	pc.People[i], pc.People[j] = pc.People[j], pc.People[i]
}

// Less implements sort.Interface
// Compares elements at indices i and j based on the selected sort field
// Thread-safe: Uses RLock for concurrent read access
// Parameters:
//   - i: first element index
//   - j: second element index
// Returns:
//   - bool: true if element i should come before element j
func (pc *PersonCollection) Less(i, j int) bool {
	pc.RLock()
	defer pc.RUnlock()

	// Determine sort direction multiplier
	// For descending order, we invert the comparison result
	multiplier := 1
	if !pc.Ascending {
		multiplier = -1
	}

	// Compare based on selected field
	// Each case handles both ascending and descending order
	switch pc.SortField {
	case "name":
		if multiplier == 1 {
			return pc.People[i].Name < pc.People[j].Name
		}
		return pc.People[i].Name > pc.People[j].Name

	case "age":
		if multiplier == 1 {
			return pc.People[i].Age < pc.People[j].Age
		}
		return pc.People[i].Age > pc.People[j].Age

	case "height":
		if multiplier == 1 {
			return pc.People[i].Height < pc.People[j].Height
		}
		return pc.People[i].Height > pc.People[j].Height

	default:
		// Default to name sorting if field is invalid
		// This ensures stable behavior even with invalid input
		if multiplier == 1 {
			return pc.People[i].Name < pc.People[j].Name
		}
		return pc.People[i].Name > pc.People[j].Name
	}
}

// Sort performs the actual sorting operation
// Thread-safe: Uses sort.Sort which calls thread-safe methods
// Returns:
//   - error: nil if successful, error if sort field is invalid
func (pc *PersonCollection) Sort() error {
	// Validate sort field before attempting sort
	if !pc.isValidSortField() {
		return fmt.Errorf("invalid sort field: %s", pc.SortField)
	}
	
	// Perform the sort using the standard library
	// sort.Sort uses our Len, Less, and Swap implementations
	sort.Sort(pc)
	return nil
}

// isValidSortField validates the sort field
// Returns:
//   - bool: true if the sort field is valid
func (pc *PersonCollection) isValidSortField() bool {
	// Define valid fields in a map for O(1) lookup
	validFields := map[string]bool{
		"name":   true,
		"age":    true,
		"height": true,
	}
	return validFields[pc.SortField]
}

// Example usage:
/*
	collection := &PersonCollection{
		People: []Person{
			{Name: "Alice", Age: 30, Height: 165.5},
			{Name: "Bob", Age: 25, Height: 180.0},
		},
		SortField: "age",
		Ascending: true,
	}
	
	err := collection.Sort()
	if err != nil {
		log.Fatal(err)
	}
*/