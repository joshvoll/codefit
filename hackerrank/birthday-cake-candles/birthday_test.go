package hackerrank

import (
	"testing"
)

func TestCandleBirth(t *testing.T) {
	arr := []int32{18, 90, 90, 13, 90, 75, 90, 8, 90, 43}
	res := birthdayCakeCandles(arr)
	if res != 5 {
		t.Fatalf("want 5 got %v ", res)
	}
}

func TestCandleBirth2(t *testing.T) {
	arr := []int32{18, 90, 90, 13, 90, 75, 90, 8, 90, 43}
	res := birthdayCakeCandles2(arr)
	if res != 5 {
		t.Fatalf("want 5, got %v ", res)
	}
}
