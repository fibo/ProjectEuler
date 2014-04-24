/*
The sum of the squares of the first ten natural numbers is,

 1^2 + 2^2 + ... + 102 = 385

The square of the sum of the first ten natural numbers is,

 (1 + 2 + ... + 10)^2 = 55^2 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/
package main

import (
	"fmt"
)

func sumOfTheSquares(n int) int64 {
	var (
		sum int64 = 0
	)

	for i := 0; i <= n; i++ {
		sum += int64(i * i)
	}

	return sum
}

func squareOfTheSums(n int) int64 {
	var (
		sum int64 = 0
	)

	for i := 0; i <= n; i++ {
		sum += int64(i)

	}

	return sum * sum
}

func main() {
	var (
		diff int64
		// first n natural numbers
		n = 100
	)

	diff = squareOfTheSums(n) - sumOfTheSquares(n)

	fmt.Println(diff)
}
