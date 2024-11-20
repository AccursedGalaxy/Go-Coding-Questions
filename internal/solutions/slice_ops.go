package solutions

/*
Slice Operations with Maps

Key Concepts:
- Slice manipulation: Modifying and creating new slices efficiently
- Map usage: Using maps for O(1) lookups and counting
- Multiple return values: Returning both processed data and metadata
- Memory management: Efficient allocation and prevention of memory leaks
- Order preservation: Maintaining original sequence while removing duplicates

This implementation demonstrates:
1. Efficient counting using hash maps for O(1) lookups
2. Slice manipulation while preserving order
3. Multiple pass algorithm for clarity and maintainability
4. Pre-allocation of slices for memory efficiency
5. Edge case handling for empty inputs and boundary conditions
*/

// CleanupSlice processes a slice of integers and removes excessive duplicates
// Parameters:
//   - numbers: input slice of integers to be processed
//   - maxOccurrences: maximum number of times any integer can appear
// Returns:
//   - []int: new slice containing unique numbers (up to maxOccurrences)
//   - map[int]int: map containing numbers that exceeded maxOccurrences and their counts
func CleanupSlice(numbers []int, maxOccurrences int) ([]int, map[int]int) {
    // First pass: Count occurrences of each number
    // Create a map to store the count of each number
    // Key: the number itself, Value: how many times it appears
    counts := make(map[int]int)
    for _, num := range numbers {
        // Increment the count for each number
        // If the number doesn't exist in the map, it's automatically initialized to 0
        counts[num]++
    }

    // Second pass: Identify excessive numbers
    // Create a map to store numbers that appear more than maxOccurrences times
    // This will be returned to inform the caller which numbers were problematic
    excessive := make(map[int]int)
    for num, count := range counts {
        // If a number appears more times than allowed,
        // add it to the excessive map with its count
        if count > maxOccurrences {
            excessive[num] = count
        }
    }

    // Third pass: Build result slice while preserving order
    // Create a map to track which numbers we've already added to the result
    // This ensures we only add each number once while maintaining original order
    seen := make(map[int]bool)
    // Pre-allocate the result slice with capacity equal to input length
    // This is an optimization to prevent multiple slice growths
    result := make([]int, 0, len(numbers))

    // Iterate through original numbers in order
    for _, num := range numbers {
        // Only add the number if:
        // 1. We haven't seen it before (seen[num] is false)
        // 2. It doesn't appear more than maxOccurrences times
        if !seen[num] && counts[num] <= maxOccurrences {
            // Append number to result slice
            result = append(result, num)
            // Mark this number as seen
            seen[num] = true
        }
    }

    // Return both the cleaned slice and the map of excessive numbers
    return result, excessive
} 