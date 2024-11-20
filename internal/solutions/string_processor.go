package solutions

import (
	"errors"
	"regexp"
	"strings"
)

/*
String Pattern Processor

Key Concepts:
- Regular expressions: Complex pattern matching and extraction
- String manipulation: Efficient string processing and replacement
- Error handling: Proper error propagation and validation
- Custom types: Using structs for organized data representation
- Memory efficiency: Proper allocation and string handling

This implementation demonstrates:
1. Regular expression compilation and usage
2. Structured data extraction from strings
3. Comprehensive error handling
4. Efficient string manipulation
5. Clean code organization using custom types
*/

// Pattern represents a key-value pair extracted from the input string
// Fields:
//   - Key: identifier found between { and :
//   - Value: content found between : and }
type Pattern struct {
    Key   string // Identifier for the pattern
    Value string // Content to replace the pattern with
}

// ProcessString processes a string containing patterns and extracts them
// Parameters:
//   - input: string containing patterns in format {key:value}
// Returns:
//   - string: processed string with all patterns replaced by their values
//   - []Pattern: slice containing all extracted patterns and their values
//   - error: error if any pattern is invalid or malformed
func ProcessString(input string) (string, []Pattern, error) {
    // Compile the regular expression pattern
    // Pattern explanation:
    // \{ - matches opening brace
    // ([^:{}]+) - captures one or more characters that aren't :, {, or }
    // : - matches literal colon
    // ([^:{}]+) - captures one or more characters that aren't :, {, or }
    // \} - matches closing brace
    re := regexp.MustCompile(`\{([^:{}]+):([^:{}]+)\}`)
    
    // Find all matches in the input string
    // FindAllStringSubmatch returns:
    // - Full match at index 0
    // - First capture group at index 1 (key)
    // - Second capture group at index 2 (value)
    matches := re.FindAllStringSubmatch(input, -1)
    
    // If no matches found, return original string unchanged
    if matches == nil {
        return input, nil, nil
    }

    // Initialize slice to store extracted patterns
    // Pre-allocate with capacity equal to number of matches
    patterns := make([]Pattern, 0, len(matches))
    
    // Keep track of the processed string
    result := input

    // Process each match
    for _, match := range matches {
        // Validate match structure
        // Each match should have 3 elements:
        // - Full match
        // - Key capture group
        // - Value capture group
        if len(match) != 3 {
            return "", nil, errors.New("invalid pattern format")
        }

        // Extract and clean key and value
        // TrimSpace removes any leading or trailing whitespace
        key := strings.TrimSpace(match[1])
        value := strings.TrimSpace(match[2])

        // Validate key and value are not empty
        if key == "" || value == "" {
            return "", nil, errors.New("empty key or value not allowed")
        }

        // Create and append new Pattern to our slice
        patterns = append(patterns, Pattern{
            Key:   key,
            Value: value,
        })

        // Replace the original pattern with its value
        // Only replace first occurrence to handle potential duplicates correctly
        result = strings.Replace(result, match[0], value, 1)
    }

    // Return processed string, extracted patterns, and nil error
    return result, patterns, nil
} 