package hackerrank

import "testing"

func TestPickingNumber(t *testing.T) {
	// t.Fatal("not implemented")
	a := []int32{4, 6, 5, 3, 3, 1}
	res := pickingNumer(a)
	if res != 3 {
		t.Fatalf("want 3, got %v ", res)
	}
}
