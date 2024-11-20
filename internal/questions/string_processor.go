package questions

/*
Challenge 0.5: String Pattern Processor

Problem: Create a string processor that can:
1. Find all patterns matching a specific format (e.g., {key:value})
2. Replace patterns with processed content
3. Validate pattern syntax
4. Return both processed string and extracted patterns

Example:
Input: "Hello {name:John}, your ID is {id:123}"
Output:
- Processed: "Hello John, your ID is 123"
- Patterns: map[string]string{"name": "John", "id": "123"}
*/

type Pattern struct {
    Key   string
    Value string
}

func ProcessString(input string) (string, []Pattern, error) {
    return "", nil, nil // Placeholder
}
