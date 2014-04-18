// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func isPalindrome(n int) bool {
	fmt.Println(string(n))
	fmt.Println(n)
	return true
}

func main() {
	i := 100 * 100
	largestPalindrome := 0

	cento := strconv.FormatInt(100, 10)
	fmt.Println(cento)
	otnec := sort.Reverse(cento)
	fmt.Println(otnec)

	//for i < 999*999 {
	for i < 99 {
		if isPalindrome(i) {
			largestPalindrome = i
		}
		i++
	}

	fmt.Println(largestPalindrome)
}
