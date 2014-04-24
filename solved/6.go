// The sum of the squares of the first ten natural numbers is,
//
// 1^2 + 2^2 + ... + 102 = 385
//
// The square of the sum of the first ten natural numbers is,
//
// (1 + 2 + ... + 10)^2 = 55^2 = 3025
//
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
//
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
package main

import (
	"fmt"
)

func sumOfTheSquares(n int) int64 {
	var (
		sum int64 = 0
	)

	return sum
}

func squareOfTheSums(n int) int64 {
	var (
		sum int64 = 0
	)

	return sum
}

func main() {
	var (
		diff int64
		// first n natural numbers
		n = 10
	)

	diff = sumOfTheSquares(n) - squareOfTheSums(n)

	fmt.Println(diff)
}
