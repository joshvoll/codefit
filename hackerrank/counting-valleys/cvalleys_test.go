package hackerrank

import "testing"

func TestCountingValleys(t *testing.T) {
	res := countingValleys(8, "UDDDUDUU")
	if res != 1 {
		t.Fatalf("Error. want=1, got=%v ", res)
	}
}
