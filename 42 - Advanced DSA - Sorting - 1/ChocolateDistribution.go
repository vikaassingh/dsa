package dsa

// Problem Description

// Given an array A of N integers where the i-th element represent the number of chocolates in the i-th packet.

// There are B number of students, the task is to distribute chocolate packets following below conditions:

// 1. Each student gets one packets.
// 2. The difference between the number of chocolates given to any two students is minimum.
// Return the minimum difference (that can be achieved) between the student who gets minimum number of chocolates and the student who gets maximum number of chocolates.

// Note: If you can't give each student 1 packet, return 0.

// Problem Constraints

// 0 <= N <= 10^5
// 1 <= A[i] <= 10^7
// 0 <= B <= 10^5

// Input Format

// The first argument contains an integer array A.

// The second argument contains an integer B.

// Output Format

// Return the minimum difference (that can be achieved) between maximum and minimum number of chocolates distributed.

// Example Input

// Input:

//   A : [3, 4, 1, 9, 56, 7, 9, 12]
//   B : 5

// Example Output

// Output:

//   6

// Example Explanation

// Explanation:

//   We can choose the packets with chocolates = [3, 4, 9, 7, 9]
//   The difference between maximum and minimum is 9-3 = 6

// Expected Output
// Provide sample input and click run to see the correct output for the provided input. Use this to improve your problem understanding and test edge cases
// Arg 1: An Integer Array, For e.g [1,2,3]
import (
	"math"
	"sort"
)

func solve(A []int, B int) int {
	sort.Ints(A)
	n := len(A)
	if B == 0 || B > n || n == 0 {
		return 0
	}
	min := A[B-1] - A[0]
	s := 1
	e := B

	for e < n {
		min = int(math.Min(float64(A[e]-A[s]), float64(min)))
		s++
		e++
	}

	return min
}
