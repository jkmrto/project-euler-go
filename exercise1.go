package main

import (
	utils "./utils"
	"fmt"
)

func main() {
	sum := 0
	limit := 1000
	dividers := []int{3, 5}

	for i := 0; i < limit; i++ {
		if utils.IsDivisibleByAny(i, dividers) {
			sum += i
		}
	}
	fmt.Print(sum, "\n")
}
