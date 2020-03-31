package kata

import "testing"

func Test_LVF(t *testing.T) {
	res := WordToMask2("family")
	if res != 66 {
		t.Fatalf("want 100, got %v ", res)
	}
}
