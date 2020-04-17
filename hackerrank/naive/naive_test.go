package hackerrank

import "testing"

func TestNaive(t *testing.T) {
	a, b := 17, 6
	res := naive(a, b)
	if res != 102 {
		t.Fatalf("want 102, got %v ", res)
	}
}
