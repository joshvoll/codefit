package hackerrank

// simpleArraySum going to sum array in a normal way
func simpleArraySum(ar []int32) int32 {
	sum := int32(0)
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}
	return sum
}

// simpleArraySumGo going to test the golang way
func simpleArraySumGo(ar []int32) int32 {
	sum := int32(0)
	for _, v := range ar {
		sum += v
	}
	return sum
}
