// The prime factors of 13195 are 5, 7, 13 and 29.
//
//What is the largest prime factor of the number 600851475143 ?
package main

import (
	"fmt"
)

func main() {
	var (
		//num    int64 = 600851475143
		num    int64 = 121
		factor int64 = 0
		i      int64 = 2
	)

	for i < num {
		fmt.Println(num % i)

		if num%i == 0 {
			factor = i
		}

		i++
	}

	fmt.Println(factor)
}
