package hackerrank

import "testing"

func TestBirthDay(t *testing.T) {
	s := []int32{1, 2, 1, 3, 2}
	d := int32(3)
	m := int32(2)
	res := birthday(s, d, m)
	if res != 2 {
		t.Fatalf("want 2, got %v ", res)
	}

}
