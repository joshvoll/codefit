package hackerrank

import "fmt"

// staircase just make it posible
func staircase(n int32) {
	for i := int32(0); i < n; i++ {
		for j := int32(0); j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := int32(0); j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
