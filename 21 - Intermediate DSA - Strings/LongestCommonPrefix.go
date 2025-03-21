package dsa

// Problem Description

// Given the array of strings A, you need to find the longest string S, which is the prefix of ALL the strings in the array.

// The longest common prefix for a pair of strings S1 and S2 is the longest string S which is the prefix of both S1 and S2.

// Example: the longest common prefix of "abcdefgh" and "abcefgh" is "abc".

// Problem Constraints

// 0 <= sum of length of all strings <= 1000000

// Input Format

// The only argument given is an array of strings A.

// Output Format

// Return the longest common prefix of all strings in A.

// Example Input

// Input 1:

// A = ["abcdefgh", "aefghijk", "abcefgh"]
// Input 2:

// A = ["abab", "ab", "abcd"];

// Example Output

// Output 1:

// "a"
// Output 2:

// "ab"

// Example Explanation

// Explanation 1:

// Longest common prefix of all the strings is "a".
// Explanation 2:

// Longest common prefix of all the strings is "ab".

// Expected Output
// Provide sample input and click run to see the correct output for the provided input. Use this to improve your problem understanding and test edge cases
// Arg 1: A String Array, For e.g ['hello','world']

func longestCommonPrefix(A []string) string {
	n := len(A)
	freq := map[int]string{}
	freq[0] = A[0]
	for i := 1; i < n; i++ {
		freq[i] = commonString(A[i], freq[i-1])
	}

	return freq[n-1]
}

func commonString(a, b string) string {
	i := 0
	n := len(a)
	if n > len(b) {
		n = len(b)
	}
	for i < n {
		if a[i] == b[i] {
			i++
			continue
		}
		break
	}

	return a[:i]
}

// vipin
// priya
