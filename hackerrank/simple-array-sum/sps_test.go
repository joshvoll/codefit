package hackerrank

import "testing"

func Test_SimpleArraySum(t *testing.T) {
	arr := []int32{1, 2, 3, 4, 10, 11}
	res := simpleArraySum(arr)
	if res != 31 {
		t.Errorf("want 31, got %v ", res)
	}
}

func Test_SimpleArraySumGo(t *testing.T) {
	arr := []int32{1, 2, 3, 4, 10, 11}
	res := simpleArraySumGo(arr)
	if res != 31 {
		t.Errorf("want 31, got %v ", res)
	}

}
func Test_SimpleArraySumGoRecap(t *testing.T) {
	arr := []int32{1, 2, 3, 4, 10, 11}
	res := simpleArraySumGoRecap(arr)
	if res != 31 {
		t.Errorf("want 31, got %v ", res)
	}

}
