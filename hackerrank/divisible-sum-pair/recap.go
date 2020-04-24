package hackerrank

import "fmt"

func divisibleSumPairsRecap(n, k int32, ar []int32) int32 {
	res := int32(0)
	for i := int32(0); i < n; i++ {
		for j := int32(1); j < n; j++ {
			if i < j {
				sum := ar[i] + ar[j]
				if sum%k == 0 {
					res++
				}
			}
		}
	}
	fmt.Println("RECAP ", res)
	return res
}
