package main

//Decompose on prime factors
//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 ?

import (
	utils "../utils"
	"fmt"
)

func main() {
	number := 600851475143

	fmt.Printf("The evaluate number is: %d\n", number)

	factor := 2
	var slice []int

	for number > 1 {
		if utils.IsDivisible(number, factor) {
			slice = append(slice, factor)
			number = number / factor

		} else {
			if factor > 2 {
				factor = factor + 2
			} else {
				factor = factor + 1
			}
		}

	}
	fmt.Printf("Largest element were: %d ", slice[len(slice)-1])
}
