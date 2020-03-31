package kata

// Grow get an array return the rduce of the numbers.
func Grow(arr []int) int {
	count := 1
	for _, val := range arr {
		count *= val
	}
	return count
}
