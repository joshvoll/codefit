package hackerrank

import "testing"

func TestVeryBigSum(t *testing.T) {
	ar := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	res := AveryBigSum(ar)
	if res != 5000000015 {
		t.Fatalf("want 5000000015, got %v ", res)
	}
}
