// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(n int64) bool {
	var (
		// Convert number in array of string containing digits
		digits []string = strings.Split(strconv.FormatInt(n, 10), "")
		l      int      = len(digits)
		j      int      = 0
	)

	for i := 0; i < l/2; i++ {
		j = l - i - 1
		if digits[i] != digits[j] {
			return false
		}
	}

	return true
}

func main() {
	var (
		start             int64 = 100
		end               int64 = 999
		largestPalindrome int64 = 0
	)

	for i := start; i <= end; i++ {
		for j := start; j <= end; j++ {
			if isPalindrome(i * j) {
				if largestPalindrome <= i*j {
					largestPalindrome = i * j
					fmt.Println(i, j, largestPalindrome)
				}
			}
		}
	}

	fmt.Println(largestPalindrome)
}
