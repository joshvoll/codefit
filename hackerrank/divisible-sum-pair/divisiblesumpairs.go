package hackerrank

func divisibleSumPairs(n, k int32, ar []int32) int32 {
	var res int32
	for i := int32(0); i < n; i++ {
		for j := int32(0); j < n; j++ {
			if i < j {
				a := ar[i] + ar[j]
				if a%k == 0 {
					res++
				}
			}
		}
	}
	return res
}
