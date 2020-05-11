package hackerrank

import "testing"

func TestDayOfProgrammer(t *testing.T) {
	var tt int32 = 2016
	var tc string = "12.09.2016"
	res := dayOfProgrammer(tt)
	if res != tc {
		t.Fatalf("want %v, got %v : ", tc, res)
	}
}
