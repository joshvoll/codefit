package kata

import "testing"

func TestFunctionalAdditon(t *testing.T) {
	add := Add(1)
	res := add(3)
	if res != 4 {
		t.Errorf("want 4, got %v ", res)
	}
}
