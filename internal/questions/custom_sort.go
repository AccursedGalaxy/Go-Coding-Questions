package questions

/*
Challenge: Custom Sort Implementation

Problem: Implement a custom sorter for a collection of Person structs that can:
1. Sort by multiple fields (name, age, height)
2. Support both ascending and descending order
3. Handle invalid sort criteria
4. Implement sort.Interface

Requirements:
- Implement sort.Interface methods
- Support multiple sort criteria
- Thread-safe implementation
*/

type Person struct {
    Name   string
    Age    int
    Height float64
}

type PersonCollection struct {
    People    []Person
    SortField string
    Ascending bool
}

func (pc *PersonCollection) Len() int           { return 0 }
func (pc *PersonCollection) Less(i, j int) bool { return false }
func (pc *PersonCollection) Swap(i, j int)      {}
