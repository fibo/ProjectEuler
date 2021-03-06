/*
Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/
package main

import "fmt"

func main() {
	sum := 0
	swap := 0
	i1 := 1
	i2 := 2

	// isEven values occur when isEven = 0, where isEven is modulo 3
	isEven := 2

	for i1 < 4000000 {
		swap = i1
		i1 = i2
		i2 += swap

		isEven += 1
		isEven = isEven % 3
		if isEven == 0 {
			sum += i1
		}
	}

	fmt.Println(sum)
}
