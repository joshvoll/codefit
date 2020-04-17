package hackerrank

import "fmt"

// give an array of number, you need to look for square that sum the total of number base on the order
// example : given this array s = {1,2,1,3,2} and d = 3 and m = 2
// need to look for two square that the sum is 3
// [1][2][1][3][2]
// the square witn number 1 and 2 sum = 3
func birthday(s []int32, d, m int32) int32 {
	var count int32
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
