package dsa

// You are given a string S, and you have to find all the amazing substrings of S.

// An amazing Substring is one that starts with a vowel (a, e, i, o, u, A, E, I, O, U).

// Input

// Only argument given is string S.
// Output

// Return a single integer X mod 10003, here X is the number of Amazing Substrings in given the string.
// Constraints

// 1 <= length(S) <= 1e6
// S can have special characters
// Example

// Input
//     ABEC

// Output
//     6

// Explanation
//     Amazing substrings of given string are :
//     1. A
//     2. AB
//     3. ABE
//     4. ABEC
//     5. E
//     6. EC
//     here number of substrings are 6 and 6 % 10003 = 6.
// Expected Output
// Provide sample input and click run to see the correct output for the provided input. Use this to improve your problem understanding and test edge cases
// Arg 1: A single String, For e.g 'anagram'

func solve(A string) int {
	n := len(A)
	ans := 0
	mod := 10003
	for i := 0; i < n; i++ {
		if (A[i] == 'a') || (A[i] == 'e') || (A[i] == 'i') || (A[i] == 'o') || (A[i] == 'u') ||
			(A[i] == 'A') || (A[i] == 'E') || (A[i] == 'I') || (A[i] == 'O') || (A[i] == 'U') {
			ans += n - i
			ans = ans % mod
		}
	}

	return ans
}
