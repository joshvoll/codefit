package kata

import (
	"fmt"
	"testing"
)

func Test_Grow(t *testing.T) {
	numbers := []int{4, 1, 1, 1, 4}
	res := Grow(numbers)
	fmt.Println(res)
	if res != 16 {
		t.Fatalf("want 6, got %v ", res)
	}
}
