package hackerrank

import (
	"fmt"
	"sort"
)

// MiniMaxSumRecap is just recap of the functions
func MiniMaxSumRecap(arr []int32) {
	sum := int64(0)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := int32(0); i < int32(len(arr)); i++ {
		sum += int64(arr[i])
	}
	fmt.Println("recap", sum-int64(arr[4]), sum-int64(arr[0]))

}

func miniMaxSumRecap(arr []int32) {
	var total, min, max int64 = 0, 0, 0
	min = int64(arr[0])
	for i := int64(0); i < int64(len(arr)); i++ {
		total += int64(arr[i])
		if int64(arr[i]) < min {
			min = int64(arr[i])
		}
		if int64(arr[i]) > max {
			max = int64(arr[i])
		}
	}
	fmt.Println("max recap : ", total-max, total-min)
}
