package hackerrank

import "fmt"

// divisibleSumPairs take n,k, ar
// Divisible Sum Pairs Complete the divisibleSumPairs function in the editor below.
// It should return the integer count of pairs meeting the criteria.
// divisibleSumPairs has the following parameter(s):
// n: the integer length of array
// ar: an array of integers
// k: the integer to divide the pair sum by
func divisibleSumPairs(n, k int32, ar []int32) int32 {
	var res int32
	for i := int32(0); i < n; i++ {
		fmt.Println("i value : ", ar[i])
		for j := i + 1; j < n; j++ {
			fmt.Println("j value: ", ar[j])
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
