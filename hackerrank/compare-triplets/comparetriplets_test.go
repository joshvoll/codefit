package hackerrank

import (
	"fmt"
	"testing"
)

func Test_CompareTriplets(t *testing.T) {
	alice := []int32{17, 28, 30}
	bob := []int32{99, 16, 8}
	res := CompareTriplets(alice, bob)
	fmt.Println(res)
}
