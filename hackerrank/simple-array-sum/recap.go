package hackerrank

func simpleArraySumGoRecap(arr []int32) int32 {
	res := int32(0)
	for i := int32(0); i < int32(len(arr)); i++ {
		res += arr[i]
	}
	return res
}
