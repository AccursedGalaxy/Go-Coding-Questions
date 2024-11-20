package solutions

import (
	"unicode"
)

/*
Palindrome Checker

Key Concepts:
- String manipulation using runes (for Unicode support)
- Character classification with unicode package
- Two-pointer technique for efficient comparison
- Case-insensitive string handling

This implementation:
1. Handles Unicode characters properly
2. Ignores non-alphanumeric characters
3. Is case-insensitive
4. Uses memory efficiently
*/

// IsPalindrome checks if a string is a palindrome while ignoring case and non-alphanumeric characters
// Parameters:
//   - s: input string to check
// Returns:
//   - bool: true if the string is a palindrome, false otherwise
func IsPalindrome(s string) bool {
    // Create a slice to store cleaned characters
    // Using rune type to properly handle Unicode characters
    cleaned := []rune{}
    
    // First pass: clean the string
    // Iterate through each character (rune) in the input string
    for _, char := range s {
        // Only include letters and numbers in our cleaned slice
        // unicode.IsLetter checks if the character is a letter in any language
        // unicode.IsNumber checks if the character is a numeric digit
        if unicode.IsLetter(char) || unicode.IsNumber(char) {
            // Convert the character to lowercase for case-insensitive comparison
            // Append it to our cleaned slice
            cleaned = append(cleaned, unicode.ToLower(char))
        }
    }

    // Two-pointer technique for palindrome checking
    // left starts from the beginning, right starts from the end
    left, right := 0, len(cleaned)-1
    
    // Continue checking while the pointers haven't met in the middle
    for left < right {
        // If characters at left and right positions don't match,
        // the string is not a palindrome
        if cleaned[left] != cleaned[right] {
            return false
        }
        // Move pointers towards the center
        left++
        right--
    }
    
    // If we've made it through the loop, all characters matched
    // Therefore, the string is a palindrome
    return true
} 