package hackerrank

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_CompareTriplets(t *testing.T) {
	alice := []int32{17, 28, 30}
	bob := []int32{99, 16, 8}
	res := CompareTriplets(alice, bob)
	fmt.Println(res)
}

func Test_CompareTripletsRecap(t *testing.T) {
	alice := []int32{5, 6, 7}
	bob := []int32{3, 6, 10}
	tc := []int32{1, 1}
	res := CompareTripletsRecap(alice, bob)
	if !reflect.DeepEqual(res, tc) {
		t.Fatalf("want %v, got %v ", tc, res)
	}
}
