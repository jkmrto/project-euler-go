package main

//Limit 4_000_000, of fib sequence
//Sum of even valued terms

import (
	fibonnacci "./fibonacci"
	utils "./utils"
	"fmt"
)

func main() {
	limit := 4000000
	fibSequence := fibonnacci.GetSequenceUpToLimit(limit)
	sum := 0

	for _, value := range fibSequence {
		if utils.IsDivisible(value, 2) {
			sum = sum + value
		}
	}

	fmt.Print(sum)
}
