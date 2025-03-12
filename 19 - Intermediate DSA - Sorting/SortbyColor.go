package dsa

// Q1. Sort by Color
// Solved
// feature icon
// Using hints except Complete Solution is Penalty free now
// Use Hint
// Problem Description

// Given an array with N objects colored red, white, or blue, sort them so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

// We will represent the colors as,

// red -> 0
// white -> 1
// blue -> 2

// Note: Using the library sort function is not allowed.

// Problem Constraints

// 1 <= N <= 1000000
// 0 <= A[i] <= 2

// Input Format

// First and only argument of input contains an integer array A.

// Output Format

// Return an integer array in asked order

// Example Input

// Input 1 :
//     A = [0, 1, 2, 0, 1, 2]
// Input 2:

//     A = [0]

// Example Output

// Output 1:
//     [0, 0, 1, 1, 2, 2]
// Output 2:

//     [0]

// Example Explanation

// Explanation 1:
//     [0, 0, 1, 1, 2, 2] is the required order.
// Explanation 2:
//     [0] is the required order

func sortColors(A []int) []int {
	var l, m int
	h := len(A) - 1
	for m <= h {
		switch A[m] {
		case 0:
			temp := A[l]
			A[l] = A[m]
			A[m] = temp
			l++
			m++
			break
		case 1:
			m++
			break
		case 2:
			temp := A[m]
			A[m] = A[h]
			A[h] = temp
			h--
			break

		}
	}

	return A
}
