package hackerrank

import (
	"reflect"
	"testing"
)

func TestBreakingRecords(t *testing.T) {
	tc := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
	tt := []int32{4, 0}
	res := breakingRecords(tc)
	if !reflect.DeepEqual(res, tt) {
		t.Fatalf("want 4 0, got %v ", res)
	}
}
