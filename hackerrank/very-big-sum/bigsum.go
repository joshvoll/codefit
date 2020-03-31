package hackerrank

// AveryBigSum going to sum an array, some time big numbers
func AveryBigSum(ar []int64) int64 {
	var sum int64
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}
	return sum
}
