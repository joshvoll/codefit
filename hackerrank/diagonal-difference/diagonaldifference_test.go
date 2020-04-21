package hackerrank

import "testing"

func TestDiagonalDifference(t *testing.T) {
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	diff := DiagonalDifference(arr)
	if diff != 15 {
		t.Errorf("want 15, got %v ", diff)
	}
}

func TestDiagonalDifferenceRecap(t *testing.T) {
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	diff := DiagonalDifferenceRecap(arr)
	if diff != 15 {
		t.Errorf("want 15, got %v ", diff)
	}
}
