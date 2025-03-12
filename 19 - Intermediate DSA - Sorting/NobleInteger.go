package dsa

import "sort"

// Problem Description

// Given an integer array A, find if an integer p exists in the array such that the number of integers greater than p in the array equals p.

// Problem Constraints

// 1 <= |A| <= 2*105
// -108 <= A[i] <= 108

// Input Format

// First and only argument is an integer array A.

// Output Format

// Return 1 if any such integer p is present else, return -1.

// Example Input

// Input 1:

//  A = [3, 2, 1, 3]
// Input 2:

//  A = [1, 1, 3, 3]

// Example Output

// Output 1:

//  1
// Output 2:

//  -1

// Example Explanation

// Explanation 1:

//  For integer 2, there are 2 greater elements in the array..
// Explanation 2:

//  There exist no integer satisfying the required conditions.

// Expected Output
// Provide sample input and click run to see the correct output for the provided input. Use this to improve your problem understanding and test edge cases
// Arg 1: An Integer Array, For e.g [1,2,3]
func solve(A []int) int {
	n := len(A)

	/*for i := 0; i < n-1; i++ {
	    minIndex := i
	    for j := i+1; j < n; j++ {
	        if A[j] < A[minIndex] {
	            minIndex = j
	        }
	    }

	    if minIndex != i {
	        tmp := A[i]
	        A[i] = A[minIndex]
	        A[minIndex] = tmp
	    }
	}*/

	sort.Ints(A)
	for i := 0; i < n; i++ {
		if n-1-i == A[i] && ((i < n-1 && A[i] != A[i+1]) || i == n-1) {
			return 1
		}
	}

	return -1
}
