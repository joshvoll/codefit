package kata

// WordToMask get a word and return the equivalent in sum
func WordToMask(s string) int {
	count := 1
	letterwithvalue := make(map[string]int)
	total := 0
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for i, v := range alphabet {
		letterwithvalue[v] += count + i
	}

	for _, j := range s {
		total += letterwithvalue[string(j)]
	}

	return total
}

// WordToMask2 jus ta test
func WordToMask2(s string) int {
	var result int
	for _, v := range s {
		result += int((v - 96))
	}
	return result
}
