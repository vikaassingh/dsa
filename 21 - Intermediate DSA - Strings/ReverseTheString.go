package dsa

// Problem Description

// You are given a string A of size N.

// Return the string A after reversing the string word by word.

// NOTE:

// A sequence of non-space characters constitutes a word.
// Your reversed string should not contain leading or trailing spaces, even if it is present in the input string.
// If there are multiple spaces between words, reduce them to a single space in the reversed string.

// Problem Constraints

// 1 <= N <= 3 * 105

// Input Format

// The only argument given is string A.

// Output Format

// Return the string A after reversing the string word by word.

// Example Input

// Input 1:
// A = "the sky is blue"
// Input 2:
// A = "this is ib"

// Example Output

// Output 1:
// "blue is sky the"
// Output 2:
// "ib is this"

// Example Explanation

// Explanation 1:
// We reverse the string word by word so the string becomes "blue is sky the".
// Explanation 2:
// We reverse the string word by word so the string becomes "ib is this".

// Expected Output
// Provide sample input and click run to see the correct output for the provided input. Use this to improve your problem understanding and test edge cases
// Arg 1: A single String, For e.g 'anagram'

/**
 * @input A : String
 *
 * @Output string.
 */
import (
	"strings"
)

func solve(A string) string {
	//A = strings.TrimSpace(A)
	byteA := []byte(A)
	byteA = reverse(byteA, 0, len(byteA)-1)

	l := 0
	for i := 0; i < len(byteA); i++ {
		if string(byteA[i]) == " " {
			byteA = reverse(byteA, l, i-1)
			l = i + 1
			// To remove extra space
			j := l
			for j < len(byteA) && string(byteA[j]) == " " {
				j++
			}
			if l != j {
				byteA = append(byteA[:l], byteA[j:]...)
			}
		} else if i == len(byteA)-1 {
			byteA = reverse(byteA, l, i)
		}
	}

	return strings.TrimSpace(string(byteA))
}

func reverse(byteA []byte, l int, r int) []byte {
	e := r
	for s := l; s <= e; s++ {
		temp := byteA[s]
		byteA[s] = byteA[e]
		byteA[e] = temp
		e--
	}
	return byteA
}
