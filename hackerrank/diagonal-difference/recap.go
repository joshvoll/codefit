package hackerrank

// DiagonalDifferenceRecap just a recap method
func DiagonalDifferenceRecap(arr [][]int32) int32 {
	lsum := int32(0)
	rsum := int32(0)
	lens := int32(len(arr))
	for i := int32(0); i < lens; i++ {
		for j := int32(0); j < lens; j++ {
			if i == j {
				lsum += arr[i][j]
			}
			if i+j == lens-1 {
				rsum += arr[i][j]
			}
		}
	}
	return rsum - lsum
}
