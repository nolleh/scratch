package demo

import "testing"

func Test_IsOdd(t *testing.T) {
	t.Parallel()

	if !IsOdd(1) {
		t.Fatal("1 should be odd")
	}

	if IsOdd(2) {
		t.Fatal("2 should not be odd")
	}
}

