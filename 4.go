// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.
package main

import "fmt"

func isPalindrome() {
	return True
}

func main() {
	i := 100 * 100
	largestPalindrome := 0

	for i < 999*999 {
		if isPalindrome(i) {
			largestPalindrome = i
		}
	}

	fmt.Println(largestPalindrome)
}
