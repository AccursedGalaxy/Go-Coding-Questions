package fractorial

/*
Challenge 1: Basic Algorithm

Problem: Write a function to calculate the factorial of a given positive integer `n`.

Input: An integer `n` (0 ≤ n ≤ 20).
Output: Return the factorial of `n`.

Key Concepts:
- Factorial (n!) is the product of all positive integers up to `n`.
- Factorial is defined as:
  - 0! = 1 (special case).
  - n! = 1 × 2 × 3 × ... × n.
- This function uses an iterative approach.
*/

func Factorial(n int) int {
    // Define the function 'Factorial' that takes an integer 'n' as input and returns an integer as output.
    // Input: n, the number for which we calculate the factorial.
    // Output: The factorial of n (n!).

    // Initialize the result variable to 1.
    // Why 1? Because multiplying by 1 doesn’t change the value, and factorial involves multiplication.
    result := 1 // At this point, result = 1.

    // Start a loop from 1 to n (inclusive).
    // This loop iterates through all integers between 1 and n, updating the result variable.
    for i := 1; i <= n; i++ {
        // Multiply the current result by the loop variable 'i' (current number in the sequence).
        // Example: For n=4, this will compute result = 1 * 2 * 3 * 4.
        result *= i // Update result by multiplying it with i. Example: result = result * i.
    }

    // Return the calculated factorial value.
    // At this point, result contains the final factorial value for the input 'n'.
    return result // Output the result to wherever the function was called.
}
