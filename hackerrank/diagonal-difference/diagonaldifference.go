package hackerrank

import (
	"math"
)

// DiagonalDifference calculate the difference for a square of matrix
// Given a square matrix,
// calculate the absolute difference between the sums of its diagonals.
// for exmaple
// 1 2 3
// 4 5 6
// 9 8 9
// The left-to-right diagonal = 1 + 5 + 8 = 15.
// The right to left diagonal = 3 + 5 + 9 = 17.
// Their absolute difference is = [15-17] = 2
func DiagonalDifference(arr [][]int32) int32 {
	lsum := int32(0)
	rsum := int32(0)
	lens := len(arr)
	for i := 0; i < lens; i++ {
		for j := 0; j < lens; j++ {
			if i == j {
				lsum += arr[i][j]
			}
			if i+j == lens-1 {
				rsum += arr[i][j]
			}
		}
	}
	dif := math.Abs(float64(lsum - rsum))
	return int32(dif)
}
