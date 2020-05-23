package hackerrank

import "fmt"

// SockMerchantRecap just practices
func SockMerchantRecap(n int32, ar []int32) int32 {
	res := int32(0)
	for i := int32(0); i < n; i++ {
		for j := int32(0); j < n-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
	fmt.Println(ar)

	return res
}
