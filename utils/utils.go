package utils

import (
	"strconv"
)

// IsDivisible: Is dividend disible between divider?
func IsDivisible(dividend int, divider int) bool {
	return (dividend%divider == 0)
}

func IsDivisibleByAny(dividend int, dividers []int) bool {
	for _, divider := range dividers {
		if IsDivisible(dividend, divider) {
			return true
		}
	}
	return false
}

func IsPalidrome(number int) bool {
	s := strconv.Itoa(number)
	length := len(s)
	//var numberSliceNormal []byte
	var numberSlice = make([]byte, length)
	// Copy from a String to Byte Slice
	copy(numberSlice, s)

	// var numberSliceReverse []rune
	for i := 0; i < length/2; i++ {
		if numberSlice[i] != numberSlice[(length-1)-i] {
			return false
		}
	}

	return true
}
