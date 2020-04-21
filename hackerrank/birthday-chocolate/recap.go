package hackerrank

// recap of the birthday
func birthdayRecap(s []int32, d, m int32) int32 {
	var res int32
	for i := int32(0); i < int32(len(s))-1; i++ {
		sum := int32(0)
		for j := int32(0); j < m; j++ {
			sum = sum + s[i+j]
		}
		if sum == d {
			res++
		}
	}
	return res
}
