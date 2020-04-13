package hackerrank

import (
	"fmt"
	"testing"
)

func TestKangaroo(t *testing.T) {
	x1, v1, x2, v2 := int32(0), int32(3), int32(4), int32(2)
	res := kangaroo(x1, v1, x2, v2)
	fmt.Println(res)
	if res != "YES" {
		t.Fatalf("want NO, got %v ", res)
	}
}
