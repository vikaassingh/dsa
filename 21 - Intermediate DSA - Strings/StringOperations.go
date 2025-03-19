package dsa

// Problem Description

// Akash likes playing with strings. One day he thought of applying following operations on the string in the given order:

// Concatenate the string with itself.
// Delete all the uppercase letters.
// Replace each vowel with '#'.
// You are given a string A of size N consisting of lowercase and uppercase alphabets. Return the resultant string after applying the above operations.

// NOTE: 'a' , 'e' , 'i' , 'o' , 'u' are defined as vowels.

// Problem Constraints

// 1<=N<=100000

// Input Format

// First argument is a string A of size N.

// Output Format

// Return the resultant string.

// Example Input

// Input 1:
// A="aeiOUz"
// Input 2:
// A="AbcaZeoB"

// Example Output

// Output 1:
// "###z###z"
// Output 2:
// "bc###bc###"

// Example Explanation

// Explanatino 1:
// First concatenate the string with itself so string A becomes "aeiOUzaeiOUz".
// Delete all the uppercase letters so string A becomes "aeizaeiz".
// Now replace vowel with '#', A becomes "###z###z".
// Explanatino 2:
// First concatenate the string with itself so string A becomes "AbcaZeoBAbcaZeoB".
// Delete all the uppercase letters so string A becomes "bcaeobcaeo".
// Now replace vowel with '#', A becomes "bc###bc###".

// Expected Output
// Provide sample input and click run to see the correct output for the provided input. Use this to improve your problem understanding and test edge cases
// Arg 1: A single String, For e.g 'anagram'

func solve(A string) string {
	concateString := A + A
	bytes := []rune{}

	for _, b := range concateString {
		if b >= 'A' && b <= 'Z' {
			continue
		}
		if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' {
			b = '#'
		}
		bytes = append(bytes, b)
	}

	return string(bytes)
}
