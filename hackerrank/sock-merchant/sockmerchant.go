package hackerrank

// SockMerchant gets number of socks with a array of color
// the solution is to order the array first
// after order the array we need to compare each number and increment res variable
// after compare on we go 2 index forwar to compare another 2 numbers
func SockMerchant(n int32, ar []int32) int32 {
	res := int32(0)
	for i := int32(0); i < n; i++ {
		for j := int32(0); j < n-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
	for i := int32(0); i < n-1; i++ {
		if ar[i] == ar[i+1] {
			res++
			i++
		}
	}

	return res
}
