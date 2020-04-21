package hackerrank

import (
	"reflect"
	"testing"
)

func TestBreakingRecords(t *testing.T) {
	tc := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
	tt := []int32{2, 4}
	res := breakingRecords(tc)
	if !reflect.DeepEqual(res, tt) {
		t.Fatalf("want 2 4, got %v ", res)
	}
}

func TestBreakingRecordsRecap(t *testing.T) {
	tc := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
	tt := []int32{2, 4}
	res := breakingRecordsRecap(tc)
	if !reflect.DeepEqual(res, tt) {
		t.Fatalf("want 2 4,  got %v ", res)
	}
}
