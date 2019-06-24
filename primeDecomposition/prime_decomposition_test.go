package primedecomposition

import (
	"testing"
)

func testDecomposeByFactor(t *testing.T) {
	rest, times := decomposeByFactor(8, 2)

	if rest != 1 {
		t.Errorf("rest was incorrect, got: %d, want: %d.", rest, 1)
	}
	if times != 3 {
		t.Errorf("times was incorrect, got: %d, want: %d.", times, 3)
	}
}

func testDecomposeByFactor144(t *testing.T) {
	rest, times := decomposeByFactor(144, 3)
	eRest, eTimes := 16, 2

	if rest != 1 {
		t.Errorf("rest was incorrect, got: %d, want: %d.", rest, eRest)
	}
	if times != 3 {
		t.Errorf("times was incorrect, got: %d, want: %d.", times, eTimes)
	}
}
