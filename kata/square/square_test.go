package kata

import "testing"

func Test_SquareSum(t *testing.T) {
	numbers := []int{1, 2}
	sqr := SquareSum(numbers)
	if sqr != 5 {
		t.Fatalf("want 5 , got %v ", sqr)
	}
}

func Test_SquareSum2(t *testing.T) {
	numbers := []int{0, 3, 4, 5}
	sqr := SquareSum2(numbers)
	if sqr != 50 {
		t.Fatalf("want 50, go %v ", sqr)
	}
}
