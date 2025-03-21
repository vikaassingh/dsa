package dsa

// Problem Description

// Given two sorted integer arrays A and B, merge B and A as one sorted array and return it as an output.

// Note: A linear time complexity is expected and you should avoid use of any library function.

// Problem Constraints

// -2×109 <= A[i], B[i] <= 2×109
// 1 <= |A|, |B| <= 5×104

// Input Format

// First Argument is a 1-D array representing  A.
// Second Argument is also a 1-D array representing B.

// Output Format

// Return a 1-D vector which you got after merging A and B.

// Example Input

// Input 1:

// A = [4, 7, 9]
// B = [2, 11, 19]
// Input 2:

// A = [1]
// B = [2]

// Example Output

// Output 1:

// [2, 4, 7, 9, 11, 19]
// Output 2:

// [1, 2]

// Example Explanation

// Explanation 1:

// Merging A and B produces the output as described above.
// Explanation 2:

//  Merging A and B produces the output as described above.

// Expected Output
// Provide sample input and click run to see the correct output for the provided input. Use this to improve your problem understanding and test edge cases
// Arg 1: An Integer Array, For e.g [1,2,3]

func solve(A []int, B []int) []int {
	var i, j, k int
	var m, n int = len(A), len(B)
	C := make([]int, m+n)
	for i < m && j < n {
		if A[i] <= B[j] {
			C[k] = A[i]
			i++
		} else {
			C[k] = B[j]
			j++
		}
		k++
	}

	for i < m {
		C[k] = A[i]
		i++
		k++
	}

	for j < n {
		C[k] = B[j]
		j++
		k++
	}

	return C
}
