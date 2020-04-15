package hackerrrank

import "testing"

func TestTotalX(t *testing.T) {
	tc1, tc2 := []int32{2, 3}, []int32{2, 4}
	res := getTotalX(tc1, tc2)
	if res != 3 {
		t.Fatalf("want 3, got %v ", res)
	}
}
