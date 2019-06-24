package main

import (
	utils "../utils"
	"fmt"
)

// Find the largest palindrome made from the product of two 3-digit numbers.
func main() {
	limit := 999
	max := 0
	product := 0

	for i := limit; i > 0; i-- {
		for j := limit; j > 0; j-- {
			product = i * j
			if utils.IsPalidrome(product) && max < product {
				max = product
				break
			}
		}
	}

	fmt.Printf("The bigger palindrome is: %d\n", max)

}
