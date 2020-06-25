package hackerrank

import (
	"strings"
)

// countingValleys coding
func countingValleys(n int32, s string) int32 {
	m := make(map[string]int)
	m["U"], m["D"] = 1, -1
	valleys, cLevel, pLevel := 0, 0, 0
	for _, ch := range strings.Split(s, "") {
		pLevel = cLevel
		cLevel += m[ch]
		if cLevel == 0 && pLevel == -1 {
			valleys++
		}
	}
	return int32(valleys)
}
