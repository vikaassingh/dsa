package dsa

// Problem Description

// Given an array of positive integers A, check and return whether the array elements are consecutive or not.

// NOTE: Try this with O(1) extra space.

// Problem Constraints

// 1 <= length of the array <= 105
// 1 <= A[i] <= 109

// Input Format

// The only argument given is the integer array A.

// Output Format

// Return 1 if the array elements are consecutive else return 0.

// Example Input

// Input 1:

//  A = [3, 2, 6, 4, 5]
// Input 2:

//  A = [1, 3, 2, 5]

// Example Output

// Output 1:

//  1
// Output 2:

//  0

// Example Explanation

// Explanation 1:

//  As you can see all the elements are consecutive (i.e 2, 3, 4, 5, 6), so we return 1.
// Explanation 2:

//  Element 4 is missing, so we return 0.

// Expected Output
// Provide sample input and click run to see the correct output for the provided input. Use this to improve your problem understanding and test edge cases
// Arg 1: An Integer Array, For e.g [1,2,3]

import "sort"

func solve(A []int) int {
	a := A
	sort.Ints(a)
	n := len(a)

	for i := 1; i < n; i++ {
		if (a[i] - a[i-1]) != 1 {
			return 0
		}
	}
	return 1
}
