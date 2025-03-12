package dsa

import "sort"

// Problem Description

// Given an array A of non-negative integers, arrange them such that they form the largest number.

// Note: The result may be very large, so you need to return a string instead of an integer.

// Problem Constraints

// 1 <= len(A) <= 100000
// 0 <= A[i] <= 2*109

// Input Format

// The first argument is an array of integers.

// Output Format

// Return a string representing the largest number.

// Example Input

// Input 1:

//  A = [3, 30, 34, 5, 9]
// Input 2:

//  A = [2, 3, 9, 0]

// Example Output

// Output 1:

//  "9534330"
// Output 2:

//  "9320"

// Example Explanation

// Explanation 1:

// Reorder the numbers to [9, 5, 34, 3, 30] to form the largest number.
// Explanation 2:

// Reorder the numbers to [9, 3, 2, 0] to form the largest number 9320.

// Expected Output
// Provide sample input and click run to see the correct output for the provided input. Use this to improve your problem understanding and test edge cases
// Arg 1: An Integer Array, For e.g [1,2,3]

func largestNumber(A []int) string {
	ans := ""
	var arr []string
	for i := 0; i < len(A); i++ {
		arr = append(arr, string(A[i]))
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i]+arr[j] > arr[j]+arr[i]
	})

	for i := 0; i < len(A); i++ {
		ans += arr[i]
	}

	pos := 0
	for ans[pos] == '0' && pos+1 < len(ans) {
		pos++
	}
	return ans
}
