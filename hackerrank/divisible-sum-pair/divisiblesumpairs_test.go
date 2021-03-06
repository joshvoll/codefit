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

func TestDivisibleSumOfPair2(t *testing.T) {
	n := int32(6)
	k := int32(3) // divisable number
	arr := []int32{1, 3, 2, 6, 1, 2}
	res := divisibleSumPairs2(n, k, arr)
	if res != 5 {
		t.Fatalf("want 5, got %v ", res)
	}

}

func TestDivisibleSumOfPairRecap(t *testing.T) {
	n := int32(6)
	k := int32(3) // divisable number
	arr := []int32{1, 3, 2, 6, 1, 2}
	res := divisibleSumPairsRecap(n, k, arr)
	if res != 5 {
		t.Fatalf("want 5, got %v ", res)
	}

}
