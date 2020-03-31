package kata

// Add return cuntions
func Add(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
