package dsa

// Problem Description

// Given a string A of size N, find and return the longest palindromic substring in A.

// Substring of string A is A[i...j] where 0 <= i <= j < len(A)

// Palindrome string:
// A string which reads the same backwards. More formally, A is palindrome if reverse(A) = A.

// Incase of conflict, return the substring which occurs first ( with the least starting index).

// Problem Constraints

// 1 <= N <= 6000

// Input Format

// First and only argument is a string A.

// Output Format

// Return a string denoting the longest palindromic substring of string A.

// Example Input

// Input 1:
// A = "aaaabaaa"
// Input 2:
// A = "abba

// Example Output

// Output 1:
// "aaabaaa"
// Output 2:
// "abba"

// Example Explanation

// Explanation 1:
// We can see that longest palindromic substring is of length 7 and the string is "aaabaaa".
// Explanation 2:
// We can see that longest palindromic substring is of length 4 and the string is "abba".

// Expected Output
// Provide sample input and click run to see the correct output for the provided input. Use this to improve your problem understanding and test edge cases
// Arg 1: A single String, For e.g 'anagram'

func longestPalindrome(A string) string {
	s := 0
	e := 0
	ans := ""
	count := 1
	for i := 0; i < len(A); i++ {
		str := PalindromeSubstring(A, i, i)
		if str[1]-str[0]+1 > count {
			count = str[1] - str[0] + 1
			s = str[0]
			e = str[1]
		}

		if i < len(A)-1 {
			str = PalindromeSubstring(A, i, i+1)
			if str[1]-str[0]+1 > count {
				count = str[1] - str[0] + 1
				s = str[0]
				e = str[1]
			}
		}
	}

	for i := s; i <= e; i++ {
		ans += string(A[i])
	}
	return ans
}

func PalindromeSubstring(str string, l int, r int) [2]int {
	var strg [2]int
	for l >= 0 && r < len(str) && str[l] == str[r] {
		strg[0] = l
		strg[1] = r
		l--
		r++
	}
	return strg
}
