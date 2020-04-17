package hackerrank

import "fmt"

func naive(a, b int) int {
	if a == 0 {
		return 0
	}
	fmt.Println(a)
	return b + naive(a-1, b)
}
