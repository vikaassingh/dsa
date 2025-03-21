package dsa

// Problem Description

// Find the Bth smallest element in given array A .

// NOTE: Users should try to solve it in less than equal to B swaps.

// Problem Constraints

// 1 <= |A| <= 100000

// 1 <= B <= min(|A|, 500)

// 1 <= A[i] <= 109

// Input Format

// The first argument is an integer array A.

// The second argument is integer B.

// Output Format

// Return the Bth smallest element in given array.

// Example Input

// Input 1:

// A = [2, 1, 4, 3, 2]
// B = 3
// Input 2:

// A = [1, 2]
// B = 2

// Example Output

// Output 1:

//  2
// Output 2:

//  2

// Example Explanation

// Explanation 1:

//  3rd element after sorting is 2.
// Explanation 2:

//  2nd element after sorting is 2.

// Expected Output
// Provide sample input and click run to see the correct output for the provided input. Use this to improve your problem understanding and test edge cases
// Arg 1: An Integer Array, For e.g [1,2,3]

func kthsmallest(A []int, B int) int {
	for i := 0; i < B; i++ {
		minIndex := i
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[minIndex] {
				minIndex = j
			}
		}
		swap(A, i, minIndex)
	}

	return A[B-1]
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
