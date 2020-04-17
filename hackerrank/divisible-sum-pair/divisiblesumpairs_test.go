package hackerrank

import "testing"

func TestDivisibleSumOfPair(t *testing.T) {
	n := int32(6)
	k := int32(3) // divisable number
	arr := []int32{1, 3, 2, 6, 1, 2}
	res := divisibleSumPairs(n, k, arr)
	if res != 5 {
		t.Fatalf("want 5, got %v ", res)
	}
}
