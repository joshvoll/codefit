package hackerrank

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
		for j := i + 1; j < n; j++ {
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

// divisibleSumPairs2 attent to do the solution on a golang way
func divisibleSumPairs2(n, k int32, ar []int32) int32 {
	var count int32
	for _, v := range ar {
		for _, d := range ar {
			if v < d {
				a := v + d
				if a%k == 0 {
					count++
				}
			}
		}
	}
	return count
}
