package utils

import (
	"testing"
)

func TestSucessPath(t *testing.T) {
	number := 565

	if isPalidrome(number) {
		t.Log("\t\t Successful test")
	} else {
		t.Errorf("Number %d has not been processed as palindrome. Test failed", number)
	}
}

func TestFailedPath(t *testing.T) {
	number := 5789

	if isPalidrome(number) {
		t.Errorf("Number %d has been processed as palindrome. Test failed", number)
	} else {
		t.Log("\t\t Successful test")
	}
}
func TestTable(t *testing.T) {
	var numbers = []struct {
		number       int
		isPalindrome bool
	}{
		{44, true},
		{123, false},
		{78787, true},
		{78789, false},
	}
	t.Log("Given the need to test downloading different content.")
	{
		for _, number := range numbers {
			if isPalidrome(number.number) == number.isPalindrome {
				t.Logf("\t\tSucess getting isPalindrome for %d", number.number)
			} else {
				t.Errorf("\t\tError getting isPalindrome for %d", number.number)
			}
		}
	}
}
