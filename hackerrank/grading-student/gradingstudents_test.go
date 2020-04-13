package hackerrank

import (
	"fmt"
	"testing"
)

func TestGradingStudents(t *testing.T) {
	tt := []int32{73, 67, 38, 33}
	tc := []int32{75, 67, 40, 33}
	res := gradingStudents(tt)
	fmt.Println(res)
	for i, r := range res {
		if r != tc[i] {
			fmt.Println("RES ", r)
			fmt.Println("test cases: ", tc[i])
			t.Fatalf("want %v, got %v ", tc[i], r)
		}
	}
}
