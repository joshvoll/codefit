package kata

// SquareSum get to numbers and return the result of square
func SquareSum(numbers []int) int {
	square := make([]int, len(numbers))
	count := 0
	for i, v := range numbers {
		square[i] += v * v
	}
	for _, j := range square {
		count += j
	}
	return count
}

// SquareSum2 just improve of the for submit
func SquareSum2(numbers []int) int {
	count := 0
	for _, v := range numbers {
		count += v * v
	}
	return count
}
