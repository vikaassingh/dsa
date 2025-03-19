package dsa

// Problem Description

// You are given a string A of size N consisting of lowercase alphabets.

// You can change at most B characters in the given string to any other lowercase alphabet such that the number of distinct characters in the string is minimized.

// Find the minimum number of distinct characters in the resulting string.

// Problem Constraints

// 1 <= N <= 100000

// 0 <= B <= N

// Input Format

// The first argument is a string A.

// The second argument is an integer B.

// Output Format

// Return an integer denoting the minimum number of distinct characters in the string.

// Example Input

// A = "abcabbccd"
// B = 3

// Example Output

// 2

// Example Explanation

// We can change both 'a' and one 'd' into 'b'.So the new string becomes "bbcbbbccb".
// So the minimum number of distinct character will be 2.

// Expected Output
// Provide sample input and click run to see the correct output for the provided input. Use this to improve your problem understanding and test edge cases
// Arg 1: A single String, For e.g 'anagram'

import (
	"sort"
)

func solve(A string, B int) int {
	n := len(A)
	if B >= n {
		return 1
	}
	arr := [26]int{}
	for i := 0; i < n; i++ {
		ch := A[i]
		index := ch - 'a'
		arr[index]++
	}

	freq := arr[:]
	sort.Ints(freq)
	total := 0
	for i := 0; i < 26; i++ {
		if freq[i] > 0 {
			total++
		}
	}

	changes := 0
	for i := 0; i < 26; i++ {
		if freq[i] != 0 && (B-freq[i]) >= 0 {
			B -= freq[i]
			changes++
		}
	}

	return total - changes
}
