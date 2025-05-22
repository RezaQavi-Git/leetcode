package problems

import "strconv"

// Convert to string solution
func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	for i := 0; i < len(xStr)/2; i++ {
		if xStr[i] != xStr[len(xStr)-(i+1)] {
			return false
		}
	}
	return true
}

// Without converting solution
func isPalindromeV2(x int) bool {
	if x < 0 {
		return false
	}
	actual := x
	reversed := 0
	for x != 0 {
		reversed = reversed*10 + x%10
		x = x / 10
	}
	return actual == reversed
}