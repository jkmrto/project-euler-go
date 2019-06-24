package primedecomposition

import (
	"../utils"
)

func decomposeByFactor(number int, factor int) (int, int) {
	return doDecomposeByFactor(number, factor, 0)
}

func doDecomposeByFactor(number int, factor int, times int) (int, int) {
	if utils.IsDivisible(number, factor) {
		return doDecomposeByFactor(number/factor, factor, times+1)
	} else {
		return number, times
	}

}
