package hackerrank

import "testing"

func TestPageCount(t *testing.T) {
	// t.Fatal("not implemented")
	t1 := int32(6)
	t2 := int32(2)
	res := pageCount(t1, t2)
	if res != 1 {
		t.Fatalf("want 1, got %v ", res)
	}
}
