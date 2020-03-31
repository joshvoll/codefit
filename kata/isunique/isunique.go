package kata

// isUnique determinate if a string has all unique chars/
// this solution is base on hashtable
func isUnique(s string) bool {
	set := make(map[rune]bool)
	for _, c := range s {
		if _, ok := set[c]; ok {
			return false
		}
		set[c] = true
	}
	return true
}

// isUnique determinate if a string has all unique chars
// this solution is without any special golang structs
func isUniqueVanilla(s string) bool {
	lens := len(s)
	for i := 0; i < lens; i++ {
		for j := i + 1; j < lens; j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}
