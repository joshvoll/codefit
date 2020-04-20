package competitive

import (
	"reflect"
	"testing"
)

func TestSorting(t *testing.T) {
	tc := []int32{1, 3, 8, 2, 9, 2, 5, 6}
	tt := []int32{1, 2, 2, 3, 5, 6, 8, 9}
	res := sorting(tc)
	if !reflect.DeepEqual(res, tt) {
		t.Fatalf("wants:  %v, got : %v ", tt, res)
	}
}
