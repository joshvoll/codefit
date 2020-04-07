package hackerrank

func birthdayCakeCandles(arr []int32) int32 {
	max := int32(0)
	res := int32(0)
	for _, v := range arr {
		if max <= v {
			max = v
		}
	}
	for _, j := range arr {
		if max == j {
			res++
		}
	}
	return res
}

func birthdayCakeCandles2(arr []int32) int32 {
	res := int32(0)
	max := int32(0)
	for i := int32(0); i < int32(len(arr)); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	for j := int32(0); j < int32(len(arr)); j++ {
		if max == arr[j] {
			res++
		}
	}
	return res
}
