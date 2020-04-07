package hackerrank

import (
	"fmt"
	"sort"
)

// MiniMaxSum is going to get an array and return the minimum and maximum numbers of 5 elements on the array
// the return should be on space
func MiniMaxSum(arr []int32) {
	// sort the array
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	ff := int64(0)
	lf := int64(0)
	lens := int64(len(arr))
	for i := int64(0); i < 4; i++ {
		ff += int64(arr[i])
		if i == 1 {
			for j := int64(1); j < lens; j++ {
				lf += int64(arr[j])
			}
		}
	}
	fmt.Println(ff, lf)
}
