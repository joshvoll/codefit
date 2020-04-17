package hackerrank

import "fmt"

func birthday(s []int32, d, m int32) int32 {
	var count int32
	fmt.Println("lengh : ", int32(len(s)), " m ", m)
	fmt.Println("custmo length ", int32(len(s))-m+1)
	for i := int32(0); i < int32(len(s))-m+1; i++ {
		fmt.Println(s[i])
		sum := int32(0)
		for j := int32(0); j < m; j++ {
			sum = sum + s[i+j]
		}
		if sum == d {
			count++
		}
	}
	return count
}
