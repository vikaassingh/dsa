package dsa

import "sort"

// Problem Description

// Given an integer array A of size N. Return 1 if the array can be arranged to form an arithmetic progression, otherwise return 0.

// A sequence of numbers is called an arithmetic progression if the difference between any two consecutive elements is the same.

// Problem Constraints

// 2 <= N <= 105

// -109 <= A[i] <= 109

// Input Format

// The first and only argument is an integer array A of size N.

// Output Format

// Return 1 if the array can be rearranged to form an arithmetic progression, otherwise return 0.

// Example Input

// Input 1:

//  A = [3, 5, 1]
// Input 2:

//  A = [2, 4, 1]

// Example Output

// Output 1:

//  1
// Output 2:

//  0

// Example Explanation

// Explanation 1:

//  We can reorder the elements as [1, 3, 5] or [5, 3, 1] with differences 2 and -2 respectively, between each consecutive elements.
// Explanation 2:

//  There is no way to reorder the elements to obtain an arithmetic progression.

// Expected Output
// Provide sample input and click run to see the correct output for the provided input. Use this to improve your problem understanding and test edge cases
// Arg 1: An Integer Array, For e.g [1,2,3]

func Solve(A []int) int {
	sort.Ints(A)
	diff := A[1] - A[0]
	for i := 1; i < len(A)-1; i++ {
		if A[i+1]-A[i] != diff {
			return 0
		}
	}

	return 1
}
