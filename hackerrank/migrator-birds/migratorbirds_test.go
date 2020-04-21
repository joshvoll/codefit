package hackerrank

import "testing"

func TestMigratorBirds(t *testing.T) {
	tc := []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}
	res := migratoryBirds(tc)
	if res != 3 {
		t.Fatalf("want 3, got %v ", res)
	}
	tc1 := []int32{1, 1, 2, 2, 3}
	res1 := migratoryBirds(tc1)
	if res1 != 1 {
		t.Fatalf("want 1, got %v ", res)
	}
}
