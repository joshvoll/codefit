package hackerrank

import "testing"

func TestBlackList(t *testing.T) {
	tt := [][]int32{
		{3, 2},
		{1, 4, 1},
		{2, 2, 2},
	}
	res := blackList(tt)
	if res != 5 {
		t.Fatalf("want 5, got %v ", res)
	}
}
