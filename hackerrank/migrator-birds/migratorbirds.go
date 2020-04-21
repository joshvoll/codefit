package hackerrank

func migratoryBirds(arr []int32) int32 {
	max, occ := arr[0], int32(0)
	for i, v := range arr {
		if v > max {
			max = v
			occ = int32(i) - 1
		}
	}
	return occ
}
